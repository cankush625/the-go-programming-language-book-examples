package testing

import (
	"fmt"
	"testing"
)

// TestGetDivision tests the GetDivision function for
// all possible valid inputs
func TestGetDivision(t *testing.T) {
	var tests = []struct {
		input [2]int32
		want  int32
	}{
		{[2]int32{1, 1}, 1},
		{[2]int32{2, 1}, 2},
	}
	for _, test := range tests {
		if got, _ := GetDivision(test.input[0], test.input[1]); got != test.want {
			t.Errorf("GetDivision(%v, %v) = %v", test.input[0], test.input[1], got)
		}
	}
}

// TestInvalidDivision tests the GetDivision function for
// all possible invalid inputs
func TestInvalidDivision(t *testing.T) {
	var tests = []struct {
		input [2]int32
		want  string
	}{
		// Division by zero error
		{[2]int32{1, 0}, "Division by zero."},
	}
	for _, test := range tests {
		if _, err := GetDivision(test.input[0], test.input[1]); err.Error() != test.want {
			t.Errorf("GetDivision(%v, %v) = %v", test.input[0], test.input[1], err.Error())
		}
	}
}

// TestGetDivisionCoverage tests the GetDivision function for
// all possible valid inputs and also generates the coverage report
// Test function name should have Coverage suffix in order to
// generate the coverage report
func TestGetDivisionCoverage(t *testing.T) {
	var tests = []struct {
		input [2]int32
		want  int32
	}{
		{[2]int32{1, 1}, 1},
		{[2]int32{2, 1}, 2},
	}
	for _, test := range tests {
		if got, _ := GetDivision(test.input[0], test.input[1]); got != test.want {
			t.Errorf("GetDivision(%v, %v) = %v", test.input[0], test.input[1], got)
		}
	}
}

// BenchmarkGetDivision benchmarks the GetDivision function
// against huge number of executions
func BenchmarkGetDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDivision(45, 5)
	}
}

// ExampleGetDivision is an example function. It serves as a
// documentation function.
func ExampleGetDivision() {
	fmt.Println(GetDivision(4, 2))
	fmt.Println(GetDivision(45, 9))
	// Output:
	// 2
	// 5
}
