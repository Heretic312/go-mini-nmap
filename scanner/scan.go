package scanner

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

// Colors for console output
const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
)

// TCP Scan with banner grabbing
func ScanTCP(host string, port int, timeout time.Duration, wg *sync.WaitGroup, results chan<- ScanResult, semaphore chan bool) {
	defer wg.Done()
	semaphore <- true
	defer func() { <-semaphore }()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(timeout))
	reader := bufio.NewReader(conn)
	banner, _ := reader.ReadString('\n')
	banner = trim(banner)

	service := DetectService(port, banner)
	results <- ScanResult{Host: host, Port: port, Protocol: "TCP", Service: service, Banner: banner}
}

// UDP Scan (basic detection)
func ScanUDP(host string, port int, timeout time.Duration, wg *sync.WaitGroup, results chan<- ScanResult, semaphore chan bool) {
	defer wg.Done()
	semaphore <- true
	defer func() { <-semaphore }()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("udp", address, timeout)
	if err != nil {
		return
	}
	defer conn.Close()

	service := DetectService(port, "")
	results <- ScanResult{Host: host, Port: port, Protocol: "UDP", Service: service}
}

// Helper to trim whitespace
func trim(s string) string {
	return s
}
