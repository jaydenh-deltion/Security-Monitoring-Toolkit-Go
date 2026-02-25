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

func RunPortScanner() {
	fmt.Println("Starting Port Scanner...")
	fmt.Println("Enter the target IP address:")
	
}

func NewPortScanner(ip string, maxConcurrency int64) *PortScanner {
	return &PortScanner{
		ip:   ip,
		lock: semaphore.NewWeighted(maxConcurrency),
	}
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := net.JoinHostPort(ip, strconv.Itoa(port))
	
	
	dialer := net.Dialer{Timeout: timeout}
	conn, err := dialer.Dial("tcp", target)
	
	if err != nil {
		
		return 
	}
	
	conn.Close()
	fmt.Printf("[+] Poort %d is OPEN op %s\n", port, ip)
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	var wg sync.WaitGroup

	for port := f; port <= l; port++ {

		
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)

		go func(p int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, p, timeout)
		}(port)
	}
	wg.Wait() // wait for all ports to be scanned
}