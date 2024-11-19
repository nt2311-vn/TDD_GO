package iteration

// const repeatCount = 5

func Repeat(phrase string, times int) string {
	var result string

	for i := 0; i < times; i++ {
		result += phrase
	}

	return result
}
