package main

import (
	"fmt"
)

func main() {
	a, b := 2, 5

	total := sum(a, b)

	defer fmt.Println(total)

	fmt.Println("la-la-land")
}

func sum(a, b int) int {
	return a + b
}
