package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.Getfloat()
	if err != nil {
		log.Fatal(err)
	}
	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f Degree Celcius", celcius)
}
