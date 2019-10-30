// Loop - printing ascii
//We can use format printing to print out the
//decimal value with %d
//hex value with %#x
//unicode code point with %#U
//a tab with \t
//a new line with \n
package main

import (
	"fmt"
)
func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i ,i)
	}
}
