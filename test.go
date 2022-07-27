package main

import (
	"context"
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	func() {
		select {
		case <-time.After(d):
		}
	}()
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		Sleep(time.Second * 5)
		cancel()
	}()

	go func() {
		for {
			c1 <- "from 1"
			Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-ctx.Done():
				fmt.Println("Time's up")
				c3 <- true
				return
			}
		}
	}()

	<-c3
	fmt.Println("Finished")
}
