package main

import "fmt"

func average(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func main() {
	var numbers = []float64{20, 10, 30}

	fmt.Printf("Average: %v\n", average(numbers))
}
