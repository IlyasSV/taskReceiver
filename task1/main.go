package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 3: Изменение статуса заказа
//Создайте структуру Order с полями ID и Status.
//Напишите метод UpdateStatus с получателем по указателю, который изменяет статус заказа на переданный.

func main() {

	customerStatus := model.Order{ID: 1, Status: 0}
	fmt.Printf("Cumstomr status before: %d\n", customerStatus)

	customerStatus.UpdateStatus(2)

	fmt.Println("После обновления: ", customerStatus)

}
