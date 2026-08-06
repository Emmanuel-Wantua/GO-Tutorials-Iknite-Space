package main

import (
	"fmt"
)

type BankAccount struct {
	currentBalance float64
}

func (a *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	a.currentBalance += float64(amount)
	return nil
}

func (a *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	if a.currentBalance < float64(amount) {
		return fmt.Errorf("you cannot withdraw %.2f when your current balance is %.2f", float64(amount), a.currentBalance)
	}
	a.currentBalance -= float64(amount)
	return nil
}

func (a BankAccount) Balance() float64 {
	return float64(a.currentBalance)
}

type SavingsAccount struct {
	BankAccount
}

func main() {
	account := SavingsAccount{
		BankAccount: BankAccount{
			currentBalance: 0,
		},
	}

	fmt.Printf("account's initial balance: %.2f\n", account.Balance())

	err := account.Deposit(1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("account's balance after initial deposit: %.2f\n", account.Balance())

	err = account.Deposit(2000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("account's balance after second deposit: %.2f\n", account.Balance())

	err = account.Withdraw(5000)
	if err != nil {
		fmt.Println(err)
	}

	err = account.Withdraw(1500)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("account's balance after first withdrawal: %.2f\n", account.Balance())

	fmt.Printf("account's current balance: %.2f\n", account.Balance())

}
