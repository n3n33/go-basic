package main

import "fmt"

func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 18 { // create koreanAge variable for using if-else things
		return false
	}

	return true
}

func canIDrink2(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(18))
}
