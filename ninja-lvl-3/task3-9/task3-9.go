// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
package main

import (
	"fmt"
)

func main() {

	favSport := "soccer"

	switch favSport {
	case "hockey":
		fmt.Println("Hockey!")
	case "soccer":
		fmt.Println("Where is the ice?")
	}
}
