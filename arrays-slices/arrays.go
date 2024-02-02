package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlices(numbers ...[]int) []int {
	sums := []int{}
	for _, slice := range numbers {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	tails := []int{}
	for _, slice := range numbers {
		if len(slice) >= 1 {
			tails = append(tails, Sum(slice[1:]))
		} else {
			tails = append(tails, 0)
		}
	}
	return tails
}
