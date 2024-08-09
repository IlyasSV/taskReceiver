package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 1: Увеличение банковского баланса
//Создайте структуру BankAccount с полем Balance.
//Напишите метод Deposit с получателем по указателю, который добавляет сумму к балансу.

func main() {
	initialBalance := model.BankAccount{Balance: 100}
	fmt.Printf("Balance before %2f\n", initialBalance.Balance)
	initialBalance.Deposit(200)
	fmt.Printf("Balance after %2f\n", initialBalance.Balance)
}
