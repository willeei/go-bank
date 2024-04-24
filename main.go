package main

import (
	"fmt"
	"go-bank/accounts"
)

func main() {
	silviaAccount := accounts.CheckingAccount{}
	silviaAccount.Holder = "Silvia"
	silviaAccount.Balance = 500.0

	fmt.Println(silviaAccount.Balance)
	fmt.Println(silviaAccount.Withdrawal(-5000))
	fmt.Println(silviaAccount.Balance)

	williamsAccount := accounts.CheckingAccount{}
	williamsAccount.Holder = "Williams"
	williamsAccount.Balance = 200.0

	fmt.Println(williamsAccount.Balance)
	fmt.Println(williamsAccount.DepositMoney(5000))
	fmt.Println(williamsAccount.Balance)

	silviaAccount.Transfer(200., &williamsAccount)
	fmt.Println(williamsAccount.Balance)
}
