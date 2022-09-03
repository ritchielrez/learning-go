package main

import "fmt"

func makeOddNumber() func() uint {
	number := uint(1)

	return func() uint {
		returnVal := number
		number += 2
		return returnVal
	}
}

func main() {
	oddNumber := makeOddNumber()

	fmt.Println(oddNumber())
	fmt.Println(oddNumber())
	fmt.Println(oddNumber())
	fmt.Println(oddNumber())
}
