package main

import (
	"fmt"
	"log"
	"golang_examples/lesson_1"
	"golang_examples/lesson_2"
	"golang_examples/lesson_3"
	"golang_examples/lesson_4"
	"golang_examples/lesson_5"
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


	sayHello := "Hello, dude!"
	responseHello := "Hello, helo... "

	lesson_5.PrintMessage(sayHello)
	fmt.Println(sayHello)

	fmt.Println(responseHello)
	lesson_5.ChangeMessage(&responseHello)
	fmt.Println(responseHello)

	num := 5
	var p *int
	p = &num

	fmt.Println(*p)
	fmt.Println(num)

	*p = 15
	fmt.Println(num)
}
