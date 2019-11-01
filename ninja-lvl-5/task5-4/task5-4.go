// Create and use an anonymous struct.
package main

import (
	"fmt"
)

func main() {

	p := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(p)
}
