package main

import (
	"fmt"
)

type Order struct {
	CustomerName string
	Dish         string
}

func Customer(name string, orders chan<- *Order) {
	menu := []string{"pizza", "pasta", "salad"}
	for _, dish := range menu {
		orders <- &Order{CustomerName: name, Dish: dish}
	}
	close(orders)
}

func Kitchen(orders <-chan *Order, dishes chan<- *Order) {
	for order := range orders {
		go func(o *Order) {
			// prepare the dish
			// ...
			// send the prepared dish to the waiter
			dishes <- o
		}(order)
	}
	close(dishes)
}

func Waiter(dishes <-chan *Order) {
	for dish := range dishes {
		fmt.Printf("Serving %s to %s\n", dish.Dish, dish.CustomerName)
	}
}

func main() {
	orders := make(chan *Order)
	dishes := make(chan *Order)

	// start multiple customers
	go Customer("Alice", orders)
	go Customer("Bob", orders)
	go Customer("Charlie", orders)

	// start the kitchen
	go Kitchen(orders, dishes)

	// start the waiter
	Waiter(dishes)
}
