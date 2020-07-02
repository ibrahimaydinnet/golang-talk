package main

import "fmt"

func main() {
	increment := incrementBy(5)
	fmt.Printf("x: %v\n", increment(10))
}

type inc func(int) int

func incrementBy(x int) inc {
	return func(y int) int {
		return x + y
	}
}
