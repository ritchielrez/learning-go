package main

import (
	"fmt"
)

func average(nums []float64) float64 {
	// Panicking is a way to interrupt a running proccess
	// like a specific function by raising a exception runtime
	// panic("Testing")

	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func biggestSalary(salaries []int) int {
	biggestSalary := salaries[0]

	for _, salary := range salaries {
		if salary > biggestSalary {
			biggestSalary = salary
		}
	}

	return biggestSalary
}

func main() {
	var numbers = []float64{20, 10, 30}

	fmt.Printf("Average: %v\n", average(numbers))

	var salaries = []int{1000, 2500, 5000, 7500, 10000}

	fmt.Printf("Biggest salary available: %v\n", biggestSalary(salaries))
}
