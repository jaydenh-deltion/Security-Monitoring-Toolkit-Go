package main // Verander naar 'systems' als dit een library is

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := net.JoinHostPort(ip, strconv.Itoa(port))
	
	// We gebruiken een Dialer voor meer controle
	dialer := net.Dialer{Timeout: timeout}
	conn, err := dialer.Dial("tcp", target)
	
	if err != nil {
		// Voor een snelle scanner printen we vaak alleen de open poorten
		return 
	}
	
	conn.Close()
	fmt.Printf("[+] %d open\n", port)
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	var wg sync.WaitGroup

	for port := f; port <= l; port++ {
		// Acquire wacht netjes tot er ruimte is, 
		// dus we hoeven niet bang te zijn voor 'too many open files'
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)

		go func(p int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, p, timeout)
		}(port)
	}
	wg.Wait() // Wacht tot alle poorten in de range klaar zijn
}

func main() {
	// 1000 is een veilige, snelle default voor de meeste systemen
	ps := &PortScanner{
		ip:   "127.0.0.1",
		lock: semaphore.NewWeighted(1000), 
	}

	fmt.Printf("Scanning %s...\n", ps.ip)
	ps.Start(1, 1024, 200*time.Millisecond)
	fmt.Println("Scan voltooid.")
}