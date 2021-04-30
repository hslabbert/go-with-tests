package main

// Sum takes a 5-element integer array and returns the sum of the elements.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes one or more int slices and returns an int slice
// of the sums of those slices' elements
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}
