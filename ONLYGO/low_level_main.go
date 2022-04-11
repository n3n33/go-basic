package main

import "fmt"

func main() {
	a := 2
	b := &a
	//a = 5
	*b = 20
	fmt.Println(a, b)
	// &a memory adder 0xc000126058
	// &b memory adder 0xc000126070
	// pointer is see through
	fmt.Println(*b) // print 2

	fmt.Println(a)

}
