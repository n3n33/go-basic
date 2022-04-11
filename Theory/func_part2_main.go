package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	// Execute code after the function returns
	// ex : close file, images, send request rest API
	defer fmt.Println("I'm done")
	//return len(name), strings.ToUpper(name)
	length = len(name) // length is not new
	uppercase = strings.ToUpper(name)
	return
}
func main() {
	totalLength, up := lenAndUpper("yena")
	// run defer after execute func lenAndUppser
	fmt.Print(totalLength, up)
}
