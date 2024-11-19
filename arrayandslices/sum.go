package arrayandslices

func Sum(numbers [5]int) (res int) {
	for _, num := range numbers {
		res += num
	}

	return res
}
