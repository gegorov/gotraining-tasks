// Create a program that shows the “if statement” in action.

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		}
		if i%2 != 0 {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
