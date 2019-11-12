// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable:
package main

import (
	"fmt"
)

func main() {
	result := closureCreator(2)
	fmt.Println(result())
}

func closureCreator(y int) func() int {

	x := y * y

	return func() int {
		x++
		return x
	}
}
