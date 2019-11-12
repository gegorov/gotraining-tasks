// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
package main

import (
	"fmt"
)

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	x := mySlice[:5]
	y := mySlice[5:10]
	z := mySlice[2:7]
	w := mySlice[1:6]

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
}
