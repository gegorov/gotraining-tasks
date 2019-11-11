// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person
package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and i'm %v years old\n", p.first, p.last, p.age)
}

func main() {
	bond := person{
		first: "James",
		last:  "Bond",
		age:   30,
	}
	bond.speak()
}
