package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	timeout := time.After(10 * time.Second)
	count := 0

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Tick! - %d\n", count)
			count++
		case <-timeout:
			fmt.Println("timeout")
			ticker.Stop()
			return
		}
	}
}
