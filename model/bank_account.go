package model

type BankAccount struct {
	Balance float32
}

func (a *BankAccount) Deposit(amount float32) {
	a.Balance += amount
}
