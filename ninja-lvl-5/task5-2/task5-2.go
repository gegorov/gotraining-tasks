// Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
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

	m := map[string]IcecreamShopCustomer{
		customer1.lastName: customer1,
		customer2.lastName: customer2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for i, w := range v.flavors {
			fmt.Printf("\t%v: %v\n", i, w)
		}

	}

}
