package main

import "fmt"

// START OMIT
func main() {
	p := person{"ibrahim", 29}

	p.speak()

	p.incrementAge()

	fmt.Println(p.age)
}

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("my name is:", p.name)
}

func (p *person) incrementAge() {
	p.age++
}

//END OMIT
