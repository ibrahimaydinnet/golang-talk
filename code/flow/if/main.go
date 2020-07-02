package main

import "fmt"

// START OMIT
func main() {
	x := 5

	if x < 10 {
		goto multiply
	} else {
		goto add
	}

multiply:
	x *= x
	fmt.Println(x)
	return

add:
	x += x
	fmt.Println(x)
	return

}

// END OMIT
