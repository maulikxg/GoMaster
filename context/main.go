package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//ctx1 := context.Background() // The root context
	//ctx2 := context.TODO()       // placholder when unsure
	//
	//fmt.Println(ctx1, ctx2)

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)

	defer cancel() // resources free

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("Your Trip is on. Lesssgoooo!")
	case <-ctx.Done():
		fmt.Println("Yout Trip is cancelled due to :", ctx.Err())
	}

}
