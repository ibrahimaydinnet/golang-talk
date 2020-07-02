package main

import "fmt"

func main() {
	items := []int{1, 2, 3, 4, 5}
	sum := sumAll(items...)
	fmt.Printf("sum of items: %v\n", sum)

}

func sumAll(items ...int) (sum int) {
	for _, item := range items {
		sum += item
	}
	return
}
