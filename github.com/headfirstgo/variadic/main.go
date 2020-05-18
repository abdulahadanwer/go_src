package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum(44, 2, 56, 665566565))
	fmt.Println(maximum(44, 2, 56, 665566565, 11111111))
	fmt.Println(inRange(44, 245445, 56, 66, 566, 565))
	fmt.Println(inRange(44, 25778786, 56, 34, 111564, 45, 11111))
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64)[]float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}
