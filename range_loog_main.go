package main

import "fmt"

func superAdd(numbers ...int) int {
	// range gives the index
	// loops with range
	for index, number := range numbers {
		fmt.Println(index, number)
	} //left column is index, right column is number
	return 1
}

func superAdd2(numbers2 ...int) int {
	for i := 0; i < len(numbers2); i++ {
		fmt.Println(numbers2[i])
	}
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
	superAdd2(1, 2, 3, 4, 5, 6)
}
