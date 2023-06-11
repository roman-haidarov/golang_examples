package lesson_4

func Increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
