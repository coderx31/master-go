package main

import "testing"

func BenchmarkCreateSlice(b *testing.B) {
	numbers := getNumbers()
	for i := 0; i < b.N; i++ {
		CreateSlice(numbers)
	}
}

func getNumbers() []int {
	numbers := make([]int, 0, 1000)
	for i := 1; i <= 1000; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
