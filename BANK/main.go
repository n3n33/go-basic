package main

import (
	"fmt"

	"github.com/n3n33/golanggolang/BANK/banking"
)

func main() {
	account := banking.Account{Owner: "yena", Balance: 1000}
	fmt.Println(account)
}
