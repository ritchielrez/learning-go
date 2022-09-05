package main

import "fmt"

func main() {
	// Defering a function runs it after all the
	// non-deferred works are done
	defer func() {
		str := recover()
		fmt.Println(str)
	}() // Automatically call the function after definition

	// Panicking makes non-defered function called later
	// to not run. Only defered functions will run
	panic("PANIC!")

	func() {
		fmt.Println("Non-deferred function ran")
	}()
}
