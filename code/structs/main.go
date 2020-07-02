package main

import "fmt"

// START OMIT
func main() {
	u := user{
		person: person{
			name: "ibrahim",
			age:  29,
		},
		password: "secret",
	}

	u.name = "halil"

	fmt.Println(u.name)
}

type person struct {
	name string
	age  int
}

type user struct {
	person
	password string
}

// END OMIT
