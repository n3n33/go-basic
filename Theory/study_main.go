package main

import "fmt"

func main() {
	//fmt.Println("Hello World!")
	const name string = "Constant" // now no type
	// Go is type language
	// Go try to guess the type
	// you need to tell go What is what

	// name = "Try to Change"
	// connat assign to name
	// name is Constans
	fmt.Println(name)

	//var varname string = "Variable" --> looks like boring
	// Go is going to guess type for you
	// Short only in function , only variable
	varname := "Variable"
	// Variables can change
	varname = "Try to change"
	fmt.Println(varname)
}

//Main for compile

// Constant is Variable but can not change. -> immutable
// Varable is Variable. Can change
