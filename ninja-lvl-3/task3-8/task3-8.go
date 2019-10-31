// Create a program that uses a switch statement with no switch expression specified.
package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("It was false")
	case true:
		fmt.Println("It was true")
	}
}
