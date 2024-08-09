package model

import "fmt"

//Задача 6: Создание нового пользователя
//Создайте структуру User с полями Username и Email. Напишите функцию CreateUser,
//которая принимает указатель на User и заполняет его поля на основе входных данных.

type User struct {
	Username string
	Email    string
}

func CreateUser(user *User, name string, email string) {
	user.Username = name
	user.Email = email
	fmt.Printf("Name: %s, email: %s\n", user.Username, user.Email)
}
