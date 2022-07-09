package main

import (
	"fmt"
	"os"
)

func main() {
	// args[] from 0th element to 1st
	args := os.Args[0:2]
	fmt.Println(args)

	// print the args[] from 2nd to last element(2nd in this case)
	for _, arg := range args[2:] {
		fmt.Println(arg)
	}
}
