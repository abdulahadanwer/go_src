package main

import (
	"dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("My name is Khan")
	fmt.Println("Your appointment is in", days, "days.")
	fmt.Println("with a follow up in", days+dates.DaysInweek, "days")
}
