package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 9: Добавление новой книги в библиотеку
//Создайте структуру Library с полем Books, которое является срезом строк.
//Напишите функцию AddBook, которая принимает указатель на Library и добавляет новую книгу в срез.

func main() {
	lib := &model.Library{Books: []string{"12", "14"}}
	fmt.Println(lib)
	model.AddBook(lib, "AAA")

	fmt.Printf("Add book: %+v\n", lib.Books)

	fmt.Println(lib)

}
