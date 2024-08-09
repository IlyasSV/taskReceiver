package model

//Задача 8: Увеличение возраста
//Создайте структуру Person с полями Name и Age.
//Напишите функцию IncreaseAge, которая принимает указатель на Person и
//увеличивает возраст на заданное количество лет.

type Person struct {
	Name string
	Age  int
}

func IncreaseAge(p *Person, age int) {
	p.Age += age
}
