package arraysandslices

func Sum(numbers []int) (result int) {
	// for i := 0; i < len(numbers); i++ {
	// 	result += numbers[i]
	// }
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}
	return
}

func SumAllTails(tailsToSum ...[]int) (result []int) {
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
			continue
		}
		result = append(result, Sum(numbers[1:]))
	}
	return
}
