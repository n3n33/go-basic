package main

import "fmt"

func main() {
	nico := map[string]string{"name": "nico", "age": "12"} // key-value
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}

	for key, _ := range nico {
		fmt.Println(key)
	}
}
