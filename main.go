package main

import (
    "time"
    "security-scanner-go/systems"
)

func main() {
    scanner := systems.NewPortScanner("127.0.0.1", 100)

    scanner.Start(1, 1000, time.Second)
}