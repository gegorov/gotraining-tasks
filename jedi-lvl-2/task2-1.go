// Write a program that prints a number in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	myNumber := 42

	fmt.Printf("%d\n", myNumber)
	fmt.Printf("%b\n", myNumber)
	fmt.Printf("%X\n", myNumber)

}
