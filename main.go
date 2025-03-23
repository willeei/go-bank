package main

import (
	"fmt"
	"go-bank/accounts"
	"go-bank/cutomer"
)

type accountVerify interface {
	Withdrawal(amount float64) string
}

func PaymentBoleto(account accountVerify, amount float64) {
	account.Withdrawal(amount)
}

func main() {
	customerWilliams := cutomer.Holder{Name: "Williams", CPF: "111.111.111-22", Profession: "Software Developer"}
	williamsAccount := accounts.SavingsAccount{
		Holder: customerWilliams, AgencyBank: 7728, AccountNumber: 111222, Operation: 2}
	williamsAccount.DepositMoney(200)

	PaymentBoleto(&williamsAccount, 130)

	fmt.Println(williamsAccount)
	fmt.Println(williamsAccount.GetBalance())
}
