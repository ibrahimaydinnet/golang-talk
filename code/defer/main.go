package main

import "fmt"

func main() {
	defer fmt.Println("and everybody")
	defer fmt.Println("gophers")

	fmt.Println("Hello")
}
