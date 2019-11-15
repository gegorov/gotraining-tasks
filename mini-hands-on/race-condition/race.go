package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)
// with race condition
// func main() {
//     fmt.Println("CPUs:", runtime.NumCPU())
//     fmt.Println("Goroutines:", runtime.NumGoroutine())
//
//     counter := 0
//
//     const gs = 100
//     var wg sync.WaitGroup
//     wg.Add(gs)
//
//     for i := 0; i < gs; i++ {
//         go func() {
//             v := counter
//             time.Sleep(time.Second)
//             runtime.Gosched()
//             v++
//             counter = v
//             wg.Done()
//         }()
//         fmt.Println("Goroutines:", runtime.NumGoroutine())
//     }
//     wg.Wait()
//     fmt.Println("Goroutines:", runtime.NumGoroutine())
//     fmt.Println("count:", counter)
// }

// without race condition
var counter int64;
func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())


    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)
            //time.Sleep(time.Second)
            runtime.Gosched()
            fmt.Println("counter:\t", atomic.LoadInt64(&counter))
            wg.Done()
        }()
        fmt.Println("Goroutines:", runtime.NumGoroutine())
    }
    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println("count:", counter)
}