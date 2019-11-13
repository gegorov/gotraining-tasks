// Starting with this code, sort the []user by
// age
// last
// Also sort each []string “Sayings” for each user
// print everything in a way that is pleasant

package main

import (
    "fmt"
    "sort"
)

type user struct {
    First   string
    Last    string
    Age     int
    Sayings []string
}

type usersByAge []user

func (u usersByAge) Len() int           { return len(u) }
func (u usersByAge) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u usersByAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type usersByLastName []user

func (u usersByLastName) Len() int           { return len(u) }
func (u usersByLastName) Less(i, j int) bool { return u[i].Last < u[j].Last }
func (u usersByLastName) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

func main() {
    u1 := user{
        First: "James",
        Last:  "Bond",
        Age:   32,
        Sayings: []string{
            "Shaken, not stirred",
            "Youth is no guarantee of innovation",
            "In his majesty's royal service",
        },
    }

    u2 := user{
        First: "Miss",
        Last:  "Moneypenny",
        Age:   27,
        Sayings: []string{
            "James, it is soo good to see you",
            "Would you like me to take care of that for you, James?",
            "I would really prefer to be a secret agent myself.",
        },
    }

    u3 := user{
        First: "M",
        Last:  "Hmmmm",
        Age:   54,
        Sayings: []string{
            "Oh, James. You didn't.",
            "Dear God, what has James done now?",
            "Can someone please tell me where James Bond is?",
        },
    }

    users := []user{u1, u2, u3}

    for i, u := range users {
        sort.Strings(u.Sayings)
        fmt.Printf("%v\t%v\t%v%v\n", i, u.Age, u.First, u.Last)
    }

    // your code goes here
    sort.Sort(usersByAge(users))
    fmt.Println("sorted by Age")

    printUsers(users)

    sort.Sort(usersByLastName(users))
    //
    fmt.Println("sorted by Last Name")
    printUsers(users)

}

func printUsers(users []user) {
    for i, u := range users {
        fmt.Printf("%v\t%v\t%v %v\n", i, u.Age, u.First, u.Last)
        for _, saying := range u.Sayings {
            fmt.Printf("\t\t%v\n", saying)
        }
    }
}
