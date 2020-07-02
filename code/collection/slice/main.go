package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

	s = make([]string, 4, 4)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

	s = append(s, "a")
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

}
