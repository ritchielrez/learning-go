package main

import "fmt"

// A function returning a function
func makeOddNumber() func() uint {
	number := uint(1)

	return func() uint {
		returnVal := number
		number += 2
		return returnVal
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func fibonacci(num int) int {
  if num == 0 {
    return 0
  } else if num == 1 {
    return 1
  }
  return fibonacci(num - 1) + fibonacci(num - 2)
}

func main() {
	oddNumber := makeOddNumber()

	for i := 1; i <= 4; i++ {
		fmt.Println(oddNumber())
	}
	fmt.Println()

	// Asigning a function to a varaible
	increment := func(num uint) uint {
		num++
		return num
	}

	for i := 1; i <= 4; i++ {
		fmt.Println(increment(oddNumber()))
	}
	fmt.Println()

	for i := 1; i <= 4; i++ {
		fmt.Println(factorial(i))
	}
  fmt.Println()

  for i := 1; i <=4; i++ {
    fmt.Println(fibonacci(i))
  }
}
