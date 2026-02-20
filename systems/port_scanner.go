package systems 

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

func NewPortScanner(ip string, maxConcurrency int64) *PortScanner {
	return &PortScanner{
		ip:   ip,
		lock: semaphore.NewWeighted(maxConcurrency),
	}
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
	fmt.Printf("[+] Poort %d is OPEN op %s\n", port, ip)
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