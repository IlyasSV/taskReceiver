package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 4: Перевод денег между счетами
//Создайте структуру Account с полями Owner и Balance.
//Напишите метод Transfer с получателем по указателю, который переводит деньги с одного счета на другой.

func main() {
	aaa := &model.Account{Owner: "AAA", Balance: 500}
	bbb := &model.Account{Owner: "BBB", Balance: 300}

	bbb.TransferTo(aaa, 300)

	fmt.Printf("AAA: %+v\n", aaa.Balance)
	fmt.Printf("BBB: %+v\n", bbb.Balance)

}
