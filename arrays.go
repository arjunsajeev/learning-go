package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(lists ...[]int) []int {
	var sums []int
	for _, list := range lists {
		sums = append(sums, Sum(list))
	}
	return sums
}
