//Using a COMPOSITE LITERAL:
//create an ARRAY which holds 5 VALUES of TYPE int
//assign VALUES to each index position.
//Range over the array and print the values out.
//Using format printing
//print out the TYPE of the array
package main

import (
	"fmt"
)

func main() {

	myArray := [5]int{}
	fmt.Println(myArray)
	for i := 0; i < len(myArray); i++ {
		myArray[i] = i * i
	}

	fmt.Println(myArray)

	for i, v := range myArray {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", myArray)

}
