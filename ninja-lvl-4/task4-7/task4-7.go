// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
//
// Range over the records, then range over the data in each record.
package main

import (
	"fmt"
)

func main() {
	bondianaCharachters := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i, v := range bondianaCharachters {
		fmt.Printf("Record: %v\n", i)
		for j, w := range v {
			fmt.Printf("\t%v\t%v\n",j, w)
		}
	}
}
