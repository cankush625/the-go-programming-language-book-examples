package main

import "fmt"

func main() {
	// Create map using map literal
	students := map[string]int{
		"alice": 28,
		"mike":  30,
	}
	fmt.Println(students) // "map[alice:28 mike:30]"

	// Create map using make function
	users := make(map[string]int)
	users["rachel"] = 32
	users["steve"] = 33
	fmt.Println(users) // "map[rachel:32 steve:33]"

	// Access map element
	fmt.Println(users["rachel"]) // "32"

	// Remove element from map
	delete(users, "steve")
	fmt.Println(users) // "map[rachel:32]"

	// Shorthand addition to map values
	users["rachel"] += 1
	fmt.Println(users["rachel"]) // "33"
	users["rachel"]++
	fmt.Println(users["rachel"]) // "34"

	// Note:
	// We cannot get the address of the map element because
	// map element is not a variable. Because growing map can
	// cause rehashing of existing elements to a new storage
	// location, thus potentially invalidating the address.

	// Check if key exist in map or not
	age, ok := users["rachel"] // age: 32, ok: true
	if !ok {
		fmt.Printf("Key does not exist. Age is %d\n", age)
	} else {
		fmt.Printf("Age is %d\n", age)
	}

	if age, ok = users["micheal"]; !ok {
		fmt.Printf("Key does not exist. Age is %d\n", age) // age: 0, ok: false
	} else {
		fmt.Printf("Age is %d\n", age)
	}
}
