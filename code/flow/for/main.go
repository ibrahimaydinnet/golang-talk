package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println("")
	i := 0
	for i < 5 {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println("")
	items := []string{"a", "b", "c", "d", "e"}
	for i, item := range items {
		fmt.Printf("%v:%v ", i, item)
	}
	fmt.Println("")
}
