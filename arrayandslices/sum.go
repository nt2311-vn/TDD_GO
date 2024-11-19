package arrayandslices

func Sum(numbers []int) (res int) {
	for _, num := range numbers {
		res += num
	}

	return res
}

func SumAll(numbers ...[]int) []int {
	res := make([]int, 0)

	for _, nums := range numbers {
		res = append(res, Sum(nums))
	}

	return res
}
