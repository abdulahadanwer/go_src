package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filepath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filepath)
		} else {
			fmt.Println(filepath)
		}
	}
	return nil
}
func main() {
	defer reportPanic()
	scanDirectory("go")
}
