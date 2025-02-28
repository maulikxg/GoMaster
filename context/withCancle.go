package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go trip(ctx) // trip started

	time.Sleep(time.Second * 3) // trip going on so have a funn

	cancel() // emergency so back to home immediatly

	time.Sleep(time.Second) // time to stop the trip

}

func trip(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("You need to stop for now . Reason :", ctx.Err())
			return
		default:
			fmt.Println("Tripping")
			time.Sleep(time.Second * 1)
		}
	}

}
