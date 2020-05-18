package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.Getfloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade > 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
