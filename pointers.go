package main

import (
	"errors"
	"fmt"
)

type Solana int

func (s Solana) String() string {
	return fmt.Sprintf("%d SOL", s)
}

type Wallet struct {
	balance Solana
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Solana) {
	w.balance += amount
}

func (w *Wallet) Balance() Solana {
	return w.balance
}

func (w *Wallet) Withdraw(amount Solana) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance = w.balance - amount
	return nil
}
