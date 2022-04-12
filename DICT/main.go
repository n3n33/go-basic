package main

import (
	"fmt"

	"github.com/n3n33/golanggolang/DICT/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	//fmt.Println(dictionary["fisrt"])
	//word := "hello"
	//definition := "greeeting"

	//err := dictionary.Add(word, definition)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//hello, _ := dictionary.Search(word)
	//fmt.Println("found ", word, "definition", hello)
	//err2 := dictionary.Add(word, definition)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	baseword := "bye"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(dictionary)
	word, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
