package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyBank    int
	accountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdrawal(amount float64) string {
	canWithdraw := amount > 0 && amount <= c.balance

	if canWithdraw {
		c.balance -= amount
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *CurrentAccount) DepositMoney(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Depósito realizado com sucesso!", c.balance
	}

	return "Valor do depósito é menor que zero!", c.balance
}

func main() {
	silviaAccount := CurrentAccount{}
	silviaAccount.holder = "Silvia"
	silviaAccount.balance = 500.0

	fmt.Println(silviaAccount.balance)
	fmt.Println(silviaAccount.Withdrawal(-5000))
	fmt.Println(silviaAccount.balance)

	williamsAccount := CurrentAccount{}
	williamsAccount.holder = "Williams"
	williamsAccount.balance = 500.0

	fmt.Println(williamsAccount.balance)
	fmt.Println(williamsAccount.DepositMoney(5000))
	fmt.Println(williamsAccount.balance)
}
