package main

import (
	"errors"
	"fmt"
)

// To make Bitcoin, just use the syntax Bitcoin(999)
type Bitcoin int

// This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string prints.
// Just added here to show it, it's already existing built-in inside fmt package
// type Stringer interface {
// 	String() string
// }

// So when we implement String, our Bitcoin type implements requirements of Stringer interface
// Then we can update our test format strings to %s so they will use String() instead
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

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

// The next requirement is for a `Withdraw` function

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}
