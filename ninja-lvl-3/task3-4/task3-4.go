// Create a for loop using this syntax
//for { }
//Have it print out the years you have been alive.

package main

import (
	"fmt"
)

func main() {
	bd := 1985
	currentYear := 2019
	for {
		if currentYear < bd {
			break
		}
		fmt.Println(currentYear)
		currentYear--
	}
}
