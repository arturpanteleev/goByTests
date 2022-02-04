package arslice

func Sum(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	res := make([]int, 0, len(slices))
	for _, s := range slices {
		res = append(res, Sum(s))
	}

	return res
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
