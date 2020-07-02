package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go loop("goroutine")

	goloop("goloop")

	loop("normal")

	time.Sleep(time.Second)
}

func goloop(from string) {
	for i := 0; i < 3; i++ {
		go fmt.Printf("%v : %v\n", from, i)
	}
}

func loop(from string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%v : %v\n", from, i)
	}
}

// END OMIT
