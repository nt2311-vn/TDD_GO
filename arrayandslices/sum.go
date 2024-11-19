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
		if len(nums) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(nums))
		}
	}

	return res
}

func SumAllTails(numbers ...[]int) []int {
	res := make([]int, 0)

	for _, nums := range numbers {
		if len(nums) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(nums[1:]))
		}
	}

	return res
}
