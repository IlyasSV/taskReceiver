package model

//Задача 10: Модификация адреса
//Создайте структуру Address с полями Street, City, Country.
//Напишите функцию UpdateAddress, которая принимает указатель на Address и изменяет адрес на новый.

type Address struct {
	Country string
	City    string
	Street  string
}

func UpdateAddress(a *Address, newCountry string, newCity string, newStreet string) {
	a.Country = newCountry
	a.City = newCity
	a.Street = newStreet
}
