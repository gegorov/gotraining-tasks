// Build and use an anonymous func
package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I'm an anonymous func!")
	}()
}
