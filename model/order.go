package model

import "fmt"

type Order struct {
	ID     int
	Status int
}

func (o *Order) UpdateStatus(status int) {
	o.Status = status
	fmt.Printf("status: %d\n", o.Status)
}
