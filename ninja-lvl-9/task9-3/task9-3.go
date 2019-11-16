// Using goroutines, create an incrementer program
// have a variable to hold the incrementer value
// launch a bunch of goroutines
// each goroutine should
// read the incrementer value
// store it in a new variable
// yield the processor with runtime.Gosched()
// increment the new variable
// write the value in the new variable back to the incrementer variable
// use waitgroups to wait for all of your goroutines to finish
// the above will create a race condition.
// Prove that it is a race condition by using the -race flag
// if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
package main

import (
    "fmt"
    "runtime"
    "sync"
)

var counter int64
var wg sync.WaitGroup

func main() {

    gr := 100
    counter = 0

    wg.Add(gr)
    fmt.Println("goroutines:", runtime.NumGoroutine())

    for i := 0; i < gr; i++ {
        go func() {
            tmp := counter
            runtime.Gosched()
            tmp++
            counter = tmp
            fmt.Println("counter:", counter)
            wg.Done()
        }()

        fmt.Println("goroutines:", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Final counter value:", counter)
    fmt.Println("goroutines:", runtime.NumGoroutine())

}
