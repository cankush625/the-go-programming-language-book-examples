package main

import "fmt"

type Car struct {
	Name string
}

// If we need to update the receiver variable; we attach
// them to the pointer type, sych as *Car
func (c *Car) GetName() string {
	return c.Name
}

func main() {
	car := &Car{"BMW"}
	fmt.Println(car.GetName()) // "BMW"

	// Note:
	// 1.
	// Either the receiver argument has the same type as the
	// receiver parameter, for example both have type T or
	// both have type *T
	// 2.
	// If the receiver argument is a variable of type T and
	// the receiver parameter has type *T. The compiler
	// implicitly takes the address of the variable.
	// 3.
	// IF the receiver argument has type *T and the receiver
	// parameter has type T, then the compiler implicitly
	// dereferences the receiver, in other words, loads the
	// value.
}
