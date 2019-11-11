//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(2, 3))
	fmt.Println(funnySum(2, 3))

}

func sum(a int, b int) int {
	return a + b
}

func funnySum(a int, b int) (int, string) {
	message := "Bang!"
	return a + b, message
}
