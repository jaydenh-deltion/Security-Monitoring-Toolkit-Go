package main

import (
    "sync"
    "security-scanner-go/systems"
)

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 100; i++ {
        wg.Add(1)
        go func(port int) {
            defer wg.Done()
            systems.ScanPort("127.0.0.1", port)
        }(i)
    }

    wg.Wait()
}