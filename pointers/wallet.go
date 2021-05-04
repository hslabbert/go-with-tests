package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is returned for insufficient funds in a Wallet.
var ErrInsufficientFunds = errors.New(":sad_trombone: you broke, mate")

// A Bitcoin is just a special, snowflakey int.
type Bitcoin int

// A Wallet holds Bitcoin. You can Deposit() bitcoin and
// retrieve a Balance().
type Wallet struct {
	balance Bitcoin
}

// Stringer is the wrapper for implementing a String method.
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit will deposit Bitcoin into a Wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw will withdraw Bitcoin from a Wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance will return the Bitcoin balance of a Wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
