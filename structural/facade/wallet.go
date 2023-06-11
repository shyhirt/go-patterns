package facade

import (
	"errors"
)

type Wallet struct {
	balance float64
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) setBalance(amount float64) {
	w.balance = amount
}
func (w *Wallet) getBalance() float64 {
	return w.balance
}

func (w *Wallet) debitBalance(amount float64) error {
	if w.balance < amount {
		return errors.New("balance is not sufficient")
	}
	w.balance = w.balance - amount
	return nil
}
