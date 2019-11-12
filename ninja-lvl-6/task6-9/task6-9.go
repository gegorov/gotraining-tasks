// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument

package main

import (
	"fmt"
)

func main() {
	exampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	processedSlice := customMap(exampleSlice, pow)

	for i, v := range processedSlice {
		fmt.Printf("%v\t%v\n", exampleSlice[i], v)
	}
}

func pow(x int) int {
	return x * x
}

func customMap(slice []int, myFunc func(x int) int) []int {

	result := make([]int, len(slice))

	for i, v := range slice {
		result[i] = myFunc(v)
	}

	return result
}
