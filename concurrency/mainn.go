package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func generateOrders(count int) []*Order {

	orders := make([]*Order, count)

	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "pending"}
	}

	return orders

}

func processOrder(orders []*Order) {

	for _, order := range orders {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		fmt.Printf("Processing order %d\n", order.ID)

	}
}

func updateOrderStatus(orders []*Order) {

	for _, order := range orders {

		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

		status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]

		order.Status = status

		fmt.Printf("Updated order %d status :  %s\n", order.ID, order.Status)

	}
}

func reportOrderStatus(orders []*Order) {

	for i := 0; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("\n---- Order Status Report ----")

		for _, order := range orders {
			fmt.Printf("Order %d : %s\n", order.ID, order.Status)
		}

		fmt.Println("---------------------------------\n")

	}

}

func main() {

	var wg sync.WaitGroup

	orders := generateOrders(10)

	wg.Add(3)

	go func() {
		defer wg.Done()
		processOrder(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatus(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrderStatus(orders)
	}()

	wg.Wait()

	fmt.Println("All the orders completed  Exiting!")

}
