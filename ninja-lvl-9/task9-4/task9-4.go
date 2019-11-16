// Fix the race condition you created in the previous exercise by using a mutex
// it makes sense to remove runtime.Gosched()
package main

import (
    "fmt"
    "runtime"
    "sync"
)

var counter int64
var wg sync.WaitGroup
var mutex = &sync.Mutex{}
func main() {

    gr := 100
    counter = 0

    wg.Add(gr)
    fmt.Println("goroutines:", runtime.NumGoroutine())

    for i := 0; i < gr; i++ {
        go func() {
            mutex.Lock()
            tmp := counter
            mutex.Unlock()
            // runtime.Gosched()
            tmp++
            mutex.Lock()
            counter = tmp

            fmt.Println("counter:", counter)
            mutex.Unlock()
            wg.Done()
        }()

        fmt.Println("goroutines:", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Final counter value:", counter)
    fmt.Println("goroutines:", runtime.NumGoroutine())

}
