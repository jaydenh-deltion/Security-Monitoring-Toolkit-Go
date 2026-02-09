package systems

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(ip string, port int) {
	address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		fmt.Printf("Port is closed\n")
		return
	}
	conn.Close()
	fmt.Printf("Port %d is open\n", port)
}