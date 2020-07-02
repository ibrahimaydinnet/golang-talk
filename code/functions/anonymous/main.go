package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = func(a, b int) (int, int) {
		return b, a
	}(a, b)

	fmt.Printf("a: %v, b: %v\n", a, b)
}
