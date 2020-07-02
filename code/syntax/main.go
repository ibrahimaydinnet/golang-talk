package main

//START OMIT
import (
	"fmt"

	rf "github.com/ibrahimaydinnet/rootfinder"
)

const precision, maxIter = 6, 100

var Min, Max int = 3, 4

func main() {
	rf := rf.New(precision, maxIter, f)
	root, iter, err := rf.Bisection(float64(Min), float64(Max))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("after %v iters, found the root: %v with Bisection method.\n", iter, root)
}

func f(x float64) float64 {
	return x*x - 10
}

// END OMIT
