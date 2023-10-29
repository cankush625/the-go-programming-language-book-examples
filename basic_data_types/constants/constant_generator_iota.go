package main

import "fmt"

func main() {
	// A const declaration may use the constant generator iota,
	// which is used to create a sequence of related values
	// without spelling out each one explicitly
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // "0 1 2 3 4 5 6"

	// Complex example using iota
	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620177411303424
		YiB // 1208925819614629174706176
	)

	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB, float64(ZiB), float64(YiB)) // "1024 1048576 1073741824 1099511627776 1125899906842624 1152921504606846976 1.1805916207174113e+21 1.2089258196146292e+24"
}
