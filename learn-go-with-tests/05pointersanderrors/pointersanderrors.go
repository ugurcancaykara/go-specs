package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// I'm making read operation and using receiver as *pointer as well, to be consistent at codebase
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
