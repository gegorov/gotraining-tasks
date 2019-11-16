// Fix the race condition you created in exercise #4 by using package atomic

package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

var counter uint64
var wg sync.WaitGroup

func main() {

    gr := 100
    counter = 0

    wg.Add(gr)
    fmt.Println("goroutines:", runtime.NumGoroutine())

    for i := 0; i < gr; i++ {
        go func() {
            atomic.AddUint64(&counter, 1)

            fmt.Println("counter:", atomic.LoadUint64(&counter))

            wg.Done()
        }()

        fmt.Println("goroutines:", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Final counter value:", counter)
    fmt.Println("goroutines:", runtime.NumGoroutine())

}
