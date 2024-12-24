package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c2 <- "two"
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 20):
			fmt.Println("timeout")
			return
		}
	}
}
