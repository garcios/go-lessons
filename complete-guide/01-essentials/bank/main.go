package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Account struct {
	Balance float64
}

// Use constants to avoid using magic number in the code - this improves readability.
// Constants for actions.
const (
	ActionGetBalance = iota + 1
	ActionDeposit
	ActionWithdraw
	ActionExit
)

const AccountFilename = "balance.txt"

func main() {
	account, err := getAccount()
	if err != nil {
		displayError(err)
		return
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println()

	for {
		showMenu()

		choice, err := getChosen()
		if err != nil {
			displayError(err)
			continue
		}

		if choice == ActionExit {
			if err := writeBalanceToFile(account.Balance); err != nil {
				displayError(err)
				return
			}

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

func getAccount() (*Account, error) {
	balance, err := readBalanceFromFile()
	if err != nil {
		return nil, err
	}

	return &Account{Balance: balance}, nil
}

func writeBalanceToFile(balance float64) error {
	balanceText := fmt.Sprintf("%f", balance)
	err := os.WriteFile(AccountFilename, []byte(balanceText), 0644)
	if err != nil {
		return err
	}

	return nil
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(AccountFilename)
	if err != nil {
		return 0, fmt.Errorf("error reading balance file: %s", err)
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing balance: %s", err)
	}
	return balance, nil
}

func displayError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

func do(action int, account *Account) error {
	switch action {
	case ActionGetBalance:
		balance, err := account.GetBalance()
		if err != nil {
			return err
		}
		fmt.Printf("You have a current balance of %.2f\n", balance)
	case ActionDeposit:
		fmt.Print("Enter amount to deposit: ")
		amount, err := scanFloat64()
		if err != nil {
			return err
		}

		err = account.Deposit(amount)
		if err != nil {
			return err
		}

		balance, err := account.GetBalance()
		if err != nil {
			return err
		}
		fmt.Printf("You have an updated balance of %.2f\n", balance)
	case ActionWithdraw:
		fmt.Print("Enter amount to withdraw: ")
		amount, err := scanFloat64()
		if err != nil {
			return err
		}

		err = account.Withdraw(amount)
		if err != nil {
			return err
		}

		balance, err := account.GetBalance()
		if err != nil {
			return err
		}

		fmt.Printf("You have an updated balance of %.2f\n", balance)
	default:
		return fmt.Errorf("invalid choice %d", action)
	}

	return nil
}

func scanFloat64() (float64, error) {
	var (
		amount float64
		err    error
	)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		amount, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("error parsing input: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning input: %s", err)
	}

	return amount, nil
}

func scanInt() (int, error) {
	var (
		number int
		err    error
	)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		number, err = strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("error parsing input: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning input: %s", err)
	}

	return number, nil
}

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

func getChosen() (int, error) {
	fmt.Print("Your choice: ")
	choice, err := scanInt()
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
