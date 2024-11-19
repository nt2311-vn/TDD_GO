package arrayandslices

func Sum(numbers []int) (res int) {
	for _, num := range numbers {
		res += num
	}

	return res
}
