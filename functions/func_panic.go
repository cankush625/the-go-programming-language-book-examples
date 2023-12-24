package main

func positive_addition(num1, num2 int) int {
	return num1 + num2
}

func MustBePositive(num int) bool {
	if num < 0 {
		panic("number should not be negative")
	}
	return true
}

func main() {
	a := positive_addition(10, 20)
	MustBePositive(a)
	b := positive_addition(-20, 10)
	MustBePositive(b)
}

// Output:
// panic: number should not be negative
//
//  goroutine 1 [running]:
//  main.MustBePositive(...)
//          /Users/ankushchavan/GolandProjects/the-go-programming-language-book-examples/functions/func_panic.go:9
//  main.main()
//          /Users/ankushchavan/GolandProjects/the-go-programming-language-book-examples/functions/func_panic.go:18 +0x30
//  exit status 2
