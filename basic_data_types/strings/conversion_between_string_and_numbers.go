package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 123

	// We can user fmt.Sprintf to convert int to string
	fmt.Sprintf("%d", num) // "123"

	// We can use function strconv.Itoa to convert int to string
	fmt.Println(strconv.Itoa(num)) // "123"

	// Format numbers in different base
	fmt.Println(strconv.FormatInt(int64(num), 2)) // "1111011"

	// Convert string representing integer to integer
	str_num := "123"
	conv_num, _ := strconv.Atoi(str_num) // conv_num is an int
	fmt.Println(conv_num)                // 123

	conv_num_64, _ := strconv.ParseInt(str_num, 10, 64) // base 10, upto 64 bits
	fmt.Println(conv_num_64)                            // 123
}
