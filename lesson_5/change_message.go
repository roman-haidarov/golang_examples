package lesson_5

func ChangeMessage(message *string) {
	*message += " (из функции ChangeMessage())"
}
