package main

import (
	"D:/04_study/golanggolang/BANK/banking"
	"fmt"
)

func main() {
	account := banking.Account{Owner: "yena", Balance: 1000}
	fmt.Println(account)
}
