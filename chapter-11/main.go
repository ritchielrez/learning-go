package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Test String"
	fmt.Printf("str = %v\n", str)
	fmt.Printf("str contains 'es': %v\n", strings.Contains(str, "es"))
	fmt.Printf("str has %v t's\n", strings.Count(str, "t"))
	fmt.Printf("str starts with 'te': %v\n", strings.HasPrefix(str, "te"))
	fmt.Printf("str ends with 'ng': %v\n", strings.HasSuffix(str, "ng"))
}
