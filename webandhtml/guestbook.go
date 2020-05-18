package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	//"text/template"
)

// Guestbook holds info about guestbook
type Guestbook struct {
	SignaturesCount int
	Signatures      []string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	checkError(scanner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	/* placeHolder := []byte("Signature list goes here")
	_, err := writer.Write([]byte(placeHolder))
	checkError(err) */
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	checkError(err)
	guestbook := Guestbook{
		SignaturesCount: len(signatures),
		Signatures:      signatures,
	}
	err = html.Execute(writer, guestbook)
	checkError(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	checkError(err)
	err = html.Execute(writer, nil)
	checkError(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	checkError(err)
	_, err = fmt.Fprintln(file, signature)
	checkError(err)
	err = file.Close()
	checkError(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	checkError(err)
	err = tmpl.Execute(os.Stdout, data)
	checkError(err)
}

func executeTemplateDemo() {
	executeTemplate("Dot is {{.}}!\n", "ABC")
	executeTemplate("Dot is {{.}}!\n", 123.5)
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", false)
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop:{{.}}\n"
	executeTemplate(templateText, []string{"rr", "aa", "bb"})
	templateText = "Prices:\n{{range .}}{{.}}\n{{end}}"
	executeTemplate(templateText, []float64{12.3, 11, 5.36})
	type Part struct {
		Name  string
		Count int
	}
	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Henry", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})
	type Subscribter struct {
		Name   string
		Rate   float64
		Active bool
	}
	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: {{.Rate}}\n{{end}}"
	subscriber := Subscribter{Name: "Abdul Ahad", Rate: 12.3, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscribter{Name: "Karen", Rate: .333, Active: false}
	executeTemplate(templateText, subscriber)
}
