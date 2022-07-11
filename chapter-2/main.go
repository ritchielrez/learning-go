package main

import "fmt"

func main() {
	// By default maps are allocated with the
	// `new` keyword. It only allocated memory,
	// doesn't initalize the map, thus will return
	// an error as a map like any other data structures
	// can't hold any value without intilization
	// var person1 map[string]string

	// person["name"] = "John"
	// person["age"] = "144"
	// person["bday"] = "9/3/1878"

	// `make` keyword initializes bothe
	person := make(map[string]string)

	person["name"] = "John"
	person["age"] = "144"
	person["bday"] = "9/3/1878"

	fmt.Println(person)
	fmt.Println()

	// Tnx to `range` keyword, we can easilly
	// iterate over keys and values of a map
	for key, value := range person {
		if key == "bday" {
			key = "birthday"
			value = "at " + value
		}
		fmt.Printf("Person's %v is %v\n", key, value)
	}

	fmt.Println()
	fmt.Println(person)
	fmt.Println()

	// `delete` keyword allow for removal
	// of key+value in a map
	delete(person, "bday")

	for key, value := range person {
		fmt.Printf("Person's %v is %v\n", key, value)
	}
}
