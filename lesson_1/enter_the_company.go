package lesson_1

import (
	"errors"
)

func EnterTheCompany(age int) (string, error) {
	if age >= 18 {
		return "Go!", nil
	}
	return "You have not been permitted to enter", errors.New("failure")
}
