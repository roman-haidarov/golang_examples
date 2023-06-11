package lesson_2

import (
	"errors"
)

func Prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Хорошего тебе начала недели!", nil
	case "вт":
		return "Желаю продуктивного дня!", nil
	case "ср":
		return "Замечательной тебе среды!", nil
	case "чт":
		return "Четверг - это маленькая пятница ;-)", nil
	case "пт":
		return "Прекрасная возможность завершить все дела до выходных!", nil
	default:	
		return "Не корректный день недели", errors.New("invalid param")
	}
}
