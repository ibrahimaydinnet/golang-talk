package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)

	m = make(map[string]int)

	m["b"] = 2
	m["d"] = 4
	m["c"] = 3
	m["a"] = 1
	fmt.Println(m)
	fmt.Println(m["a"])
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
	if v, ok := m["x"]; ok {
		fmt.Println(v)
	}
}
