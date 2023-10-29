package main

import "fmt"

func main() {
	const (
		KB = 1000      // 1024
		MB = 1000 * KB // 1000000
		GB = 1000 * MB // 1000000000
		TB = 1000 * GB // 1000000000000
		PB = 1000 * TB // 1000000000000000
		EB = 1000 * PB // 1000000000000000000
		ZB = 1000 * EB // 1000000000000000000000
		YB = 1000 * ZB // 1000000000000000000000000
	)

	// Convert results greater that int to float64 to avoid overflow
	fmt.Println(KB, MB, GB, TB, PB, EB, float64(ZB), float64(YB)) // "1000 1000000 1000000000 1000000000000 1000000000000000 1000000000000000000 1e+21 1e+24"

}
