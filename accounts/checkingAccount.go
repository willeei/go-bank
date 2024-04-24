package accounts

type CheckingAccount struct {
	Holder        string
	AgencyBank    int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) Withdrawal(amount float64) string {
	canWithdraw := amount > 0 && amount <= c.Balance

	if canWithdraw {
		c.Balance -= amount
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *CheckingAccount) DepositMoney(amount float64) (string, float64) {
	if amount > 0 {
		c.Balance += amount
		return "Depósito realizado com sucesso!", c.Balance
	}

	return "Valor do depósito é menor que zero!", c.Balance
}

func (c *CheckingAccount) Transfer(transferAmount float64, accountTarget *CheckingAccount) bool {
	if transferAmount <= c.Balance && transferAmount > 0 {
		c.Balance -= transferAmount
		accountTarget.DepositMoney(transferAmount)
		return true
	}

	return false
}
