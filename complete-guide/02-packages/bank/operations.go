package main

import "fmt"

func (a *Account) Deposit(money float64) error {
	if money <= 0 {
		return fmt.Errorf("invalid amount. Deposit amount %.2f must be greater than zero ", money)
	}

	a.Balance += money

	return nil
}

func (a *Account) Withdraw(money float64) error {
	if a.Balance < money {
		return fmt.Errorf("invalid amount. Amount %.2f must be equal or less than current balance", money)
	}

	a.Balance -= money

	return nil
}

func (a *Account) GetBalance() (float64, error) {
	return a.Balance, nil
}
