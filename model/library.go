package model

//Задача 9: Добавление новой книги в библиотеку
//Создайте структуру Library с полем Books, которое является срезом строк.
//Напишите функцию AddBook, которая принимает указатель на Library и добавляет новую книгу в срез.

type Library struct {
	Books []string
}

func AddBook(l *Library, book string) {
	l.Books = append(l.Books, book)
}
