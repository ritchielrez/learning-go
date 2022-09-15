package main

import "fmt"

func one(value *int) {
	*value = 1
}

func main() {
	value := new(int)
	fmt.Println(*value)
	one(value)
	fmt.Println(*value)
}
