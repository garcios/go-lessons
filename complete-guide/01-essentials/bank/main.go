package main

import (
	"fmt"
	"strings"
)

type Account struct {
	Balance float64
}

func main() {
	account := &Account{Balance: 1000.0}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println()

	for {
		showMenu()

		choice, err := getChosen()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = validateInput(choice)
		if err != nil {
			displayError(err)
			continue
		}

		if choice == 4 {
			fmt.Println()
			fmt.Println("Thanks for banking with us!")
			break
		}

		err = do(choice, account)
		if err != nil {
			displayError(err)
			continue
		}

	}
}

func displayError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

func validateInput(choice int) error {
	if choice < 0 || choice > 4 {
		return fmt.Errorf("invalid choice %d", choice)
	}

	return nil
}

func do(action int, account *Account) error {
	switch action {
	case 1:
		fmt.Printf("You have a current balance of %.2f\n", account.GetBalance())
	case 2:
		fmt.Print("Enter amount to deposit: ")
		var amount float64
		_, err := fmt.Scanf("%f", &amount)
		if err != nil {
			return err
		}

		account.Deposit(amount)

		fmt.Printf("You have an updated balance of %.2f\n", account.GetBalance())
	case 3:
		fmt.Print("Enter amount to withdraw: ")
		var amount float64
		_, err := fmt.Scanf("%f", &amount)
		if err != nil {
			return err
		}

		err = account.Withdraw(amount)
		if err != nil {
			return err
		}

		fmt.Printf("You have an updated balance of %.2f\n", account.GetBalance())
	}

	return nil
}

func (a *Account) Deposit(money float64) {
	a.Balance += money
}

func (a *Account) Withdraw(money float64) error {
	if a.Balance < money {
		return fmt.Errorf("your balance is lower than %.2f", money)
	}

	a.Balance -= money

	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func getChosen() (int, error) {
	var choice int
	fmt.Print("Your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return 0, fmt.Errorf("invalid choice: %s", err.Error())
	}

	return choice, nil
}

func showMenu() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
