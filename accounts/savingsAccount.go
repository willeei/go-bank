package accounts

import "go-bank/cutomer"

type SavingsAccount struct {
	Holder                               cutomer.Holder
	AgencyBank, AccountNumber, Operation int
	balance                              float64
}

func (c *SavingsAccount) Withdrawal(amount float64) string {
	canWithdraw := amount > 0 && amount <= c.balance

	if canWithdraw {
		c.balance -= amount
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *SavingsAccount) DepositMoney(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Depósito realizado com sucesso!", c.balance
	}

	return "Valor do depósito é menor que zero!", c.balance
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
