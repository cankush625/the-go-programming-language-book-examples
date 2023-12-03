package main

import "fmt"

func reverse_array(nums *[5]int) {
	for i := 0; i < len(nums)/2; i++ {
		end := len(nums) - i - 1
		nums[i], nums[end] = nums[end], nums[i]
	}
}

func main() {
	numbers := [5]int{2, 5, 3, 7, 8}
	fmt.Println(numbers) // "[2 5 3 7 8]"
	reverse_array(&numbers)
	fmt.Println(numbers) // "[8 7 3 5 2]"
}
