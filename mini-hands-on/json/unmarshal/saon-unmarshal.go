package main

import (
    "encoding/json"
    "fmt"
)

type person struct {
    First string `json:"First"`
    Last  string `json:"Last"`
    Age   int    `json:"Age"`
}

var peopleSlice []person

func main() {

    s := `[{"First":"James","Last":"Bond","Age":99},{"First":"Miss","Last":"Moneyp","Age":30}]`

    bs := []byte(s)

    fmt.Printf("%T\n", s)
    fmt.Printf("%T", bs)

    err := json.Unmarshal(bs, &peopleSlice)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(peopleSlice)

    fmt.Println("\nnr\tfirst\tlast\tage")
    for i, v := range peopleSlice {
        fmt.Printf("%v\t%v\t%v\t%v\n", i, v.First, v.Last, v.Age)
    }

}
