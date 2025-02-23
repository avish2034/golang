package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func newOrder(id string, amount float32, status string, createdAt time.Time) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: createdAt,
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// if any struct use one time so you create like this
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	myOrder := order{
		id:     "123",
		amount: 12.00,
		status: "Pending",
	}
	myOrder2 := newOrder("1234", 134, "paid", time.Now())
	myOrder.changeStatus("Accepted")
	myOrder.createdAt = time.Now()

	fmt.Println("Order Struct", myOrder, myOrder.getAmount(), myOrder2.amount, language)
}
