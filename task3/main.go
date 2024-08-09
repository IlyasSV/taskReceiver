package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 5: Уменьшение запаса товара
//Создайте структуру Product с полями Name и Stock.
//Напишите метод Sell с получателем по указателю, который уменьшает запас на переданное количество единиц.

func main() {
	product := model.Product{Name: "TTT", Stock: 100}

	product.Sell(50)

	fmt.Printf("Количество запаса товара: Name - %s, Stock: %d\n", product.Name, product.Stock)
}
