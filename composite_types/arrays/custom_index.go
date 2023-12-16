package main

import "fmt"

func main() {
	type Currency int
	const (
		USD Currency = iota
		INR
		EUR
		GTQ
	)

	// We can specify list of index and value pairs
	symbol := [...]string{USD: "$", INR: "₹", EUR: "€", GTQ: "Q"}
	fmt.Println(INR, symbol[INR]) // "1 ₹"
}
