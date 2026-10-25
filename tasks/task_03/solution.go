package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}

	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz", nil
	} else if n%3 == 0 {
		return "Fizz", nil
	} else if n%5 == 0 {
		return "Buzz", nil
	} else {
		return strconv.Itoa(n), nil
	}
}
