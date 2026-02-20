package main

import (
    "sync"
    "port_scanner/systems"
)

func main() {
    var wg sync.WaitGroup
    scanner := systems.PortScanner{
        ip: 
        "127.0.0.1",
        startPort: 1,
        endPort: 100,
    }
