package main

import "fmt"

func main() {
	defer handleError()

	fireAnError()
}

func fireAnError() {
	panic("something went wrong")
}

func handleError() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
