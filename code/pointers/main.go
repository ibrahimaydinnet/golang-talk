package main

import "fmt"

func main() {
	x := 5.
	px := &x

	fmt.Println(*px, px)

	square(px)
	fmt.Println(x)
}

func square(x *float64) {
	*x = *x * *x
}
