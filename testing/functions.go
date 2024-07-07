package testing

import "errors"

// GetDivision returns the division of two numbers
func GetDivision(num1, num2 int32) (int32, error) {
	if num2 == 0 {
		return 0, errors.New("Division by zero.")
	}
	return num1 / num2, nil
}
