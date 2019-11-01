// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
package main

import (
	"fmt"
)

func main() {
	bondCharachters := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	bondCharachters["octopussy"] = []string{`bikini`, `surfing`, `poison`}

	delete(bondCharachters, `no_dr`)

	for key, value := range bondCharachters {
		fmt.Printf("%v:\t\n", key)

		for index, v := range value {
			fmt.Printf("\t%v.\t%v\n", index, v)
		}
	}
}
