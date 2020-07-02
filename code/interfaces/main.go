package main

import "fmt"

// START OMIT
func main() {
	p := person{name: "ibrahim"}
	tryToSpeak(p)
}

func tryToSpeak(c canSpeak) {
	c.speak()
}

type canSpeak interface {
	speak()
}

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("my name is:", p.name)
}

// END OMIT
