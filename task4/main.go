package main

import "taskReceiver1/model"

//Задача 6: Создание нового пользователя
//Создайте структуру User с полями Username и Email. Напишите функцию CreateUser,
// которая принимает указатель на User и заполняет его поля на основе входных данных.

func main() {
	user := &model.User{}

	model.CreateUser(user, "Петя", "pet@gmail.com")
}
