package systems

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip string 
	lock *semaphore.Weighted
}

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func Scanport(ip string, port int, timeout time.Duration) bool {