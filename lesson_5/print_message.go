package lesson_5

import "fmt"

func PrintMessage(message string) {
	message += " (из функции PrintMessage())"
	fmt.Println(message)
}
