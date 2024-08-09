package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 8: Увеличение возраста
//Создайте структуру Person с полями Name и Age. Напишите функцию IncreaseAge,
// которая принимает указатель на Person и увеличивает возраст на заданное количество лет.

func main() {
	person := &model.Person{Name: "Joni", Age: 20}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	model.IncreaseAge(person, 10)
	fmt.Printf("Name: %s, Age после исправления: %d\n", person.Name, person.Age)
}
