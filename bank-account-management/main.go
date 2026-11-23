package main

import (
	"errors"
	"fmt"
)

type account struct {
	accountNumber string
	balance       float64
	ownerName     string
}

func (acc *account) deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	acc.balance += amount
	fmt.Printf("Deposited $%.2f to %s. New balance: $%.2f\n",
		amount, acc.accountNumber, acc.balance)

	return nil
}

func (acc *account) withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	if acc.balance < amount {
		return fmt.Errorf("insufficient funds in %s. Balance: $%.2f. Tried to withdraw: $%.2f",
			acc.accountNumber, acc.balance, amount)
	}
	acc.balance -= amount
	fmt.Printf("Withdrew $%.2f from %s. New balance: $%.2f\n",
		amount, acc.accountNumber, acc.balance)

	return nil
}

func (acc *account) getBalance() float64 {
	return acc.balance
}

func (acc *account) String() string {
	return fmt.Sprintf("Account [%s] Owner: %s, Balance: $%.2f",
		acc.accountNumber, acc.ownerName, acc.balance)
}

type savingsAccount struct {
	account
	interestRate float64
}

func (sa *savingsAccount) addInterest() {
	interest := sa.balance * sa.interestRate
	fmt.Printf("Adding interest $%.2f to savings account %s. ", interest, sa.accountNumber)
	err := sa.deposit(interest)
	if err != nil {
		fmt.Printf("AddInterest: Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type overdraftAccount struct {
	account
	overdraftLimit float64
}

func (oa *overdraftAccount) withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	if (oa.balance + oa.overdraftLimit) < amount {
		return fmt.Errorf("Withdrawal of $%.2f exeeds overdraft limit for %s. Available including overdraft: $%.2f",
			amount, oa.accountNumber, oa.balance+oa.overdraftLimit)
	}
	oa.balance -= amount
	fmt.Printf("Withdrew $%.2f from overdraft account %s. New balance: $%.2f\n",
		amount, oa.accountNumber, oa.balance)

	return nil

}

func main() {
	fmt.Println("--- Bank Account system ---")

	savAcc := savingsAccount{
		account: account{
			accountNumber: "SAV001",
			balance:       1000.00,
			ownerName:     "Alice Saver",
		},
		interestRate: 0.02,
	}

	fmt.Println("--- Saving Account Operations ---")
	fmt.Println(savAcc.account.String())

	amount := 200.00
	err := savAcc.deposit(amount)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n",
			amount, err)
	}

	savAcc.addInterest()

	amount = 50.00
	err = savAcc.withdraw(amount)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}

	fmt.Println("Final Savings Details: ", savAcc.account.String())

	ovdAcc := overdraftAccount{
		account: account{
			accountNumber: "OVD001",
			balance:       100.00,
			ownerName:     "Bob Spender",
		},
		overdraftLimit: 200.00,
	}

	fmt.Println("--- Overdraft Account Operations ---")
	fmt.Println(ovdAcc.account.String())

	amount = 50.00
	err = ovdAcc.deposit(amount)
	if err != nil {
		fmt.Println("Error depositing: ", err)
	}

	amount = 200.00
	err = ovdAcc.withdraw(amount)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}

	amount = 100.00
	err = ovdAcc.withdraw(amount)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}
	err = ovdAcc.withdraw(amount)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}

}
