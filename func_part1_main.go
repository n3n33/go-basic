package main

import (
	"fmt"
	"strings"
)

//Need to tell compiler what everything is
// making faster for you

func multiply(a, b int) int {
	return a * b
}

//functions have multiple values return
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("yena")
	// totalLength, _ := lenAndUpper("yena")
	// Underscore is ignore value
	// Compiler is ignoring that
	fmt.Println(totalLength, upperName)
	//fmt.Println(totalLength)

	repeatMe("yy", "ee", "nn", "aa")
}
