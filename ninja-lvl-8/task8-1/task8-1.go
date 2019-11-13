// marshal the []user to JSON
// type user struct {
//	first string
//	age   int
// }
//
// func main() {
//	u1 := user{
//		first: "James",
//		age:   32,
//	}
//
//	u2 := user{
//		first: "Moneypenny",
//		age:   27,
//	}
//
//	u3 := user{
//		first: "M",
//		age:   54,
//	}
//
//	users := []user{u1, u2, u3}
//
//	fmt.Println(users)
//
//	// your code goes here
// }
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    First string `json:"first"`
    Age   int    `json:"age"`
}

func main() {
    u1 := User{
        First: "James",
        Age:   32,
    }

    u2 := User{
        First: "Moneypenny",
        Age:   27,
    }

    u3 := User{
        First: "M",
        Age:   54,
    }

    users := []User{u1, u2, u3}

    fmt.Println(users)

    // your code goes here
    bs, err := json.Marshal(users)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(bs))

}
