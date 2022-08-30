package main

import (
	"fmt"
)

func main() {
	groceries := []string{"Apple", "Banana", "Orange"}
	groceries = append(groceries, "Pineapple")
	fmt.Println(groceries)

	buylist := []string{"Chewing-gum", "Dark Fantasy"}
	buylist = append(buylist, groceries...)

	for _, buylistitem := range buylist {
		fmt.Println(buylistitem)
	}
}
