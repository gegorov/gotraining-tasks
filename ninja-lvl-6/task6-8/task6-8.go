// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
package main

import (
	"fmt"
)

func main() {
	a := myFunc()
	a()
}

func myFunc() func() {
	fmt.Println("i'm outside")
	return func() {
		fmt.Println("I'm inside")
	}
}
