package main

import "fmt"

func main() {
	arr_str := [2]string{"Hello", "World!"}

	for index, value := range arr_str {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	cashbacks := []int{1000, 2000, 3000, 4000}
	biggest_cashback := cashbacks[0]

	for _, cashback := range cashbacks {
		if cashback > biggest_cashback {
			biggest_cashback = cashback
		}
	}

	fmt.Printf("\nBiggest cashback avalaible is %v dollars\n", biggest_cashback)
}
