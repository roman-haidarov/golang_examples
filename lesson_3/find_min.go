package lesson_3

import (
	"errors"
)

func FindMin(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	}

	min := numbers[0]

	for _, i := range numbers {
		if !isInt(i) {
			return 0, errors.New("non-integer value found in numbers")
		}

		if i < min {
			min = i
		}
	}

	return min, nil
}

func isInt(n interface{}) bool {
	_, ok := n.(int)
	return ok
}
