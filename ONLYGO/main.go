package main

import "fmt"

func main() {
	names := []string{"yena", "kwon", "golang"}
	fmt.Println(names)
	names = append(names, "april")
	fmt.Println(names)
}
