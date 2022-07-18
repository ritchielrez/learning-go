package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("go", "env")
	// `cmd.Output()` functions returns the output with
	// an error value
	out, _ := cmd.Output()
	// `cmd.Output()` returns the output as a bytes array
	// strings library can be used for string conversion
	// as the bytes array contains the ascii value of each
	// characters
	cmdOutput := strings.TrimSuffix(string(out), "\n")
	// cmdOutput := string(out)
	fmt.Println([]byte(cmdOutput))

	testStr := "Test String \n"
	testBytes := []byte(testStr)
	fmt.Printf("\n%v", testBytes)
	
	// The suffix(ending chars) would be removed from
	// the 1st arg provided to the `strings.TrimSuffix()`
	testStrTrim := strings.TrimSuffix(testStr, "\n")
	testBytesTrim := []byte(testStrTrim)
	fmt.Printf("\n%v", testBytesTrim)
	fmt.Println()
}
