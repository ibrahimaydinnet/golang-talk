package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	fmt.Println("ping")

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "pong 1"
	}()

	go func() {
		ch2 <- "pong 2"
	}()

	msg1 := <-ch1
	fmt.Println(msg1)

	msg2 := <-ch2
	fmt.Println(msg2)
}

// END OMIT
