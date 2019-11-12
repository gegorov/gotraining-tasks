// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
package main

import (
	"fmt"
)

func main() {
	bondCharacters := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	bondCharacters["octopussy"] = []string{`bikini`, `surfing`, `poison`}

	delete(bondCharacters, `no_dr`)

	for key, value := range bondCharacters {
		fmt.Printf("%v:\t\n", key)

		for index, v := range value {
			fmt.Printf("\t%v.\t%v\n", index, v)
		}
	}
}
