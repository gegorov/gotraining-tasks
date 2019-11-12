package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 3, 565, 2423, 234, 2}

	fmt.Println("Using range")
	for i, v := range mySlice {
		fmt.Printf("index: %v,\t value: %v\n", i, v)
	}

	fmt.Println("\nUsing old-fashioned for loop")
	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("index: %v,\t value: %v\n", i, mySlice[i])

	}
}
