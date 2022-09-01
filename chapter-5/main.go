package main

import (
	"fmt"
)

func main() {
	groceries := []string{"Apple", "Banana", "Orange"}
	groceries = append(groceries, "Pineapple")
	fmt.Println(groceries)

	var buylist = make([]string, 2)
	copy(buylist, groceries)

	for _, buylistitem := range buylist {
		fmt.Println(buylistitem)
	}
}
