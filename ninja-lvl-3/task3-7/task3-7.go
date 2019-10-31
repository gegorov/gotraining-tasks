//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else if i%2 != 0 {
			fmt.Printf("%v is odd\n", i)
		} else {
			fmt.Println("Something went wrong")
		}
	}
}
