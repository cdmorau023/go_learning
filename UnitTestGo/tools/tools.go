package tools

import "errors"

func ElementByIndex(slice []int, index int) int {
	return slice[index]
}
func ElementByIndex2(slice []int, index int) (result int, err error) {

	if index < 0 || index >= len(slice) {
		err = errors.New("index out of range")
		return
	}

	result = slice[index]
	return
}
