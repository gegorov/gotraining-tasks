package main

import (
    "encoding/json"
    "fmt"
)

type person struct {
    First string
    Last  string
    Age  int
}

func main() {

    jb := person{
        First: "James",
        Last:  "Bond",
        Age:   99,
    }

    mp := person{
        First: "Miss",
        Last:  "Moneyp",
        Age:   30,
    }

    people := []person{
        jb,
        mp,
    }

    fmt.Println(people)

    bs, err := json.Marshal(people)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%T",bs)
    fmt.Println(string(bs))

}
