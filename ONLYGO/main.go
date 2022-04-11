package main

import "fmt"

//struct are like object and flexible than map
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimch", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico)
	fmt.Println(nico.name)
}
