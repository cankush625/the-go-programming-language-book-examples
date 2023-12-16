package main

import "fmt"

type Student struct {
	Id      int
	Name    string
	Address Address
}

type Address struct {
	FirstLine string
	Landmark  string
	City      string
	State     string
	Country   string
}

func main() {
	address := Address{FirstLine: "205", City: "Pune", State: "Maharashtra"}
	student := Student{Id: 1, Name: "Rachel", Address: address}
	// Access student name and address
	// To access address, we need to access address related fields from
	// Address struct
	fmt.Printf("Name: %s\tCity: %s\tState: %s\n", student.Name, student.Address.City, student.Address.State) // "Name: Rachel    City: Pune      State: Maharashtra"
}
