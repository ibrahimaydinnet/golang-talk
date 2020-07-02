package main

import "fmt"

func main() {
	fmt.Printf("is %v root of f : %v\n", 2, checkRoot(2, f))

}

func f(x float64) float64 {
	return 3*x - 6
}

func checkRoot(x float64, f func(float64) float64) bool {
	return f(x) == 0.0
}
