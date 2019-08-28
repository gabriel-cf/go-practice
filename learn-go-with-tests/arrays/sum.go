package arrays

func Sum(numbers []int) int {
	acum := 0
	//for i := 0; i < len(numbers); i++ {
	//	acum += numbers[i]
	//}
	for _, number := range numbers {
		acum += number
	}
	return acum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, slice := range slices {
		sums[i] = Sum(slice)
	}
	return sums
}