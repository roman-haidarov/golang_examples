package main

import (
	"fmt"
	"log"
	"golang_examples/lesson_1"
	"golang_examples/lesson_2"
)

func main() {
	message, err := lesson_1.EnterTheCompany(23)
	if err != nil {
		log.Println(message, err)
		return
	}

	fmt.Println(message)

	message, err = lesson_2.Prediction("вт")
	if err != nil {
		log.Println(message, err)
		return
	}

	fmt.Println(message)
}
