package pointerserrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoint int

type Stringer interface {
	String() string
}

func (b Bitcoint) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoint
}

func (w *Wallet) Deposit(amt Bitcoint) {
	w.balance += amt
}

func (w *Wallet) Balance() Bitcoint {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoint) error {
	if w.balance >= amt {
		w.balance -= amt
		return nil
	}

	return ErrInsufficientFunds
}
