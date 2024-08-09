package model

type Account struct {
	Owner   string
	Balance float32
}

func (ac *Account) TransferTo(to *Account, amount float32) {
	ac.Balance -= amount
	to.Balance += amount
}
