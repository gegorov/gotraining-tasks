// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
package main

import (
	"fmt"
)

const (
	currentYear = 2019
	x           = iota + currentYear
	y           = iota + currentYear
	z           = iota + currentYear
	w           = iota + currentYear
)

func main() {
	fmt.Println(x, y, z, w)
}
