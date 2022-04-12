package main

import (
	"fmt"
	"log"

	"github.com/n3n33/golanggolang/BANK/banking"
)

func main() {
	//account := banking.Account{Owner: "yena"}
	account := banking.NewAccount("sumin")
	//fmt.Println(account)
	//account.Deposit(10)
	//fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
	fmt.Println(account.Balance(), account.Owner())
}
