package banking

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { //receiver
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance ofof your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amoint from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount

	return nil
}
