package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))

	}

	return sums
}

func SumAllTails(numbersToSumOfTails ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSumOfTails {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums

}
