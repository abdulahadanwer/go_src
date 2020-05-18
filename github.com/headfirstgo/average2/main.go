// average2 calculates the average of several numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

func average(numbers ...float64) float64 {
	fmt.Println(numbers)
	var sum float64 = 0.0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
	fmt.Println(float64(len(numbers)))
	return sum / float64(len(numbers))
}
