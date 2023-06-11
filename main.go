package main

import (
	"fmt"
	"log"
	"golang_examples/lesson_1"
	"golang_examples/lesson_2"
	"golang_examples/lesson_3"
	"golang_examples/lesson_4"
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


	number, err := lesson_3.FindMin(23, -34, 254)
	if err != nil {
		log.Println(number, err)
		return
	}

	fmt.Println(number)


	inc := lesson_4.Increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
