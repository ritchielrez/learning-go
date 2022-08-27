package main

import "fmt"

func main() {
    arr_str := [2]string{"Hello", "World!"}

    for index, value := range arr_str {
        fmt.Printf("Index: %v, Value: %v\n", index, value)
    }
}
