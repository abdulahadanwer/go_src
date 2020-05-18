package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Rate:", s.Rate)
	fmt.Println("Active:", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 4.69
	s.Active = true
	s.Address.Street = "5316 Mall Drive West"
	s.Address.City = "Lansing"
	s.Address.State = "Mi"
	s.Address.PostalCode = "48917"
	return &s
}

func main() {
	subscriber1 := defaultSubscriber("Adam Gilchrist")
	subscriber1.Rate = .36
	fmt.Println(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth McFarlane")
	printInfo(subscriber2)

	var subscriber3 magazine.Subscriber
	appDiscount(&subscriber3)
	fmt.Println("subscriber 3:", subscriber3)
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)

	subscriber4 := magazine.Subscriber{Name: "Mushtaq Ahmed", Rate: 6.66, Active: true}
	printInfo(&subscriber4)

	employee := magazine.Employee{Name: "Wasim Akram", Salary: 64000}
	fmt.Println(employee)

	address := magazine.Address{Street: "5316 Mall Drive West", City: "Lansing", State: "Mi", PostalCode: "48917"}
	fmt.Println(address)

	sub := magazine.Subscriber{Name: "Abdul Ahad", Rate: 6.66, Active: true}
	sub.Address.Street = "5316 Mall Drive West"
	sub.Address.City = "Lansing"
	sub.Address.State = "Mi"
	sub.Address.PostalCode = "48917"
	fmt.Println(sub)
}

func appDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}
