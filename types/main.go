package main

import (
	"fmt"
)

// Liters define litre type
type Liters float64

// Gallon define gallon type
type Gallon float64

// Milliliters define millimeters type
type Milliliters float64

//ToGallon converts liters to gallon
func (l Liters) ToGallon() Gallon {
	return Gallon(l * 0.264)
}

//ToGallon converts mililiters to gallon
func (m Milliliters) ToGallon() Gallon {
	return Gallon(m * 0.000264)
}

func main() {
	var carFuel Liters
	var busFuel Gallon

	carFuel = Liters(10.0)
	busFuel = Gallon(29.3)
	fmt.Println(carFuel, busFuel)

	truckFuel := Liters(45.6)
	fmt.Println(truckFuel)

	soda := Liters(1.2)
	fmt.Printf("%0.3f litres equal %0.3f gallons\n", soda, soda.ToGallon())

	water := Milliliters(2666664)
	fmt.Printf("%0.3f mililiters equal %0.3f gallons", water, water.ToGallon())
}
