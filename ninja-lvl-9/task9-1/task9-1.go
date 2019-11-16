// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func main() {

    wg.Add(2)
    go func() {
        fmt.Println("inside first goroutine")
        wg.Done()
    }()
    go func() {
        fmt.Println("inside second goroutine")
        wg.Done()
    }()

    wg.Wait()
    fmt.Println("Done")
}
