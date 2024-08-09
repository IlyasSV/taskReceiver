package model

//Создайте структуру Product с полями Name и Stock.
//Напишите метод Sell с получателем по указателю, который уменьшает запас на переданное количество единиц.

type Product struct {
	Name  string
	Stock int
}

func (p *Product) Sell(amount int) {
	p.Stock -= amount
}
