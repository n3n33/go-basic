package main

import (
	"fmt"

	"github.com/n3n33/golanggolang/DICT/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	fmt.Println(dictionary["fisrt"])
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
