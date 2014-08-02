package goconvey

import (
	"errors"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为 0")
	}
	return a / b, nil
}
