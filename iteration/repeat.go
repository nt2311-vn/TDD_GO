package iteration

const repeatCount = 5

func Repeat(phrase string) string {
	var result string

	for i := 0; i < repeatCount; i++ {
		result += phrase
	}

	return result
}
