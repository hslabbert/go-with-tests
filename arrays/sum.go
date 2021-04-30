package main

// Sum takes a 5-element integer array and returns the sum of the elements.
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
