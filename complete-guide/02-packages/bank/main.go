package main

import (
	"fmt"
	"github.com/oskiegarcia/bank/fileops"
	"github.com/oskiegarcia/bank/scanner"

	"github.com/Pallinder/go-randomdata"
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	account, err := getAccount()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s to Go Bank!\n", randomdata.FirstName(randomdata.Male))
	fmt.Println()

	for {
		showOptions()

		choice, err := getChosen()
		if err != nil {
			displayError(err)
			continue
		}

		if choice == ActionExit {
			if err := fileops.WriteFloatToFile(AccountFilename, account.Balance); err != nil {
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
	balance, err := fileops.ReadFloatFromFile(AccountFilename)
	if err != nil {
		return nil, err
	}

	return &Account{Balance: balance}, nil
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
		amount, err := scanner.ScanFloat64()
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
		amount, err := scanner.ScanFloat64()
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

func getChosen() (int, error) {
	fmt.Print("Your choice: ")
	choice, err := scanner.ScanInt()
	if err != nil {
		return 0, fmt.Errorf("invalid choice: %s", err.Error())
	}

	return choice, nil
}
