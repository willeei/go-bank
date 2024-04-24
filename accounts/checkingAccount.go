package accounts

import "go-bank/cutomer"

type CheckingAccount struct {
	Holder        cutomer.Holder
	AgencyBank    int
	AccountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdrawal(amount float64) string {
	canWithdraw := amount > 0 && amount <= c.balance

	if canWithdraw {
		c.balance -= amount
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *CheckingAccount) DepositMoney(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Depósito realizado com sucesso!", c.balance
	}

	return "Valor do depósito é menor que zero!", c.balance
}

func (c *CheckingAccount) Transfer(transferAmount float64, accountTarget *CheckingAccount) bool {
	if transferAmount <= c.balance && transferAmount > 0 {
		c.balance -= transferAmount
		accountTarget.DepositMoney(transferAmount)
		return true
	}

	return false
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
