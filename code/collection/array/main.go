package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)
	fmt.Printf("len: %d, cap: %d\n", len(arr), cap(arr))

	arr[3] = 3
	fmt.Println(arr)
	fmt.Printf("len: %d, cap: %d\n", len(arr), cap(arr))
}
