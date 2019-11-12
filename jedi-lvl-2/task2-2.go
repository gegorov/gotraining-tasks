// Using the following operators, write expressions and assign their values to variables:
//==
//<=
//>=
//!=
//<
//>
//Now print each of the variables.
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 42 == 42
	b := 42 <= 42
	c := 42 >= 42
	d := 42 != 45
	e := 42 < 100
	f := 42 > math.Pi

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
