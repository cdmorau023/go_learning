package internal_calculator

import "errors"

func sum(a, b int) int {
	return a + b
}

func rest(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func division(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
