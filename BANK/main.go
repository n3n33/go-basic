package main

import (
	"fmt"

	"github.com/n3n33/golanggolang/BANK/banking"
)

func main() {
	//account := banking.Account{Owner: "yena"}
	account := banking.NewAccount("sumin")
	fmt.Println(account)
}
