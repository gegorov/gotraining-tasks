// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import (
	"fmt"
)

type IcecreamShopCustomer struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {

	customer1 := IcecreamShopCustomer{
		firstName: "James",
		lastName:  "Bond",
		flavors: []string{
			"vanilla",
			"martini",
			"vodka",
		},
	}

	customer2 := IcecreamShopCustomer{
		firstName: "Money",
		lastName:  "Penny",
		flavors: []string{
			"chocolate",
			"strawberry",
			"sex on the beach",
		},
	}

	fmt.Println(customer1.firstName, customer1.lastName)
	for i, v := range customer1.flavors {
		fmt.Printf("\t%v: %v\n", i, v)
	}

	fmt.Println(customer2.firstName, customer2.lastName)
	for i, v := range customer2.flavors {
		fmt.Printf("\t%v: %v\n", i, v)
	}

}
