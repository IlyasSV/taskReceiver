package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 10: Модификация адреса
//Создайте структуру Address с полями Street, City, Country.
//Напишите функцию UpdateAddress, которая принимает указатель на Address и изменяет адрес на новый.

func main() {
	user := &model.Address{Country: "Rasha", City: "Kazan", Street: "Bayma1"}
	fmt.Println(user)
	model.UpdateAddress(user, "Japan", "Tokya", "afasd5")
	fmt.Printf("New address: %+v\n", user)

}
