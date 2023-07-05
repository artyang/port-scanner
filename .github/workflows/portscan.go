package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"

	"github.com/pschlump/godebug/parse"
)

func main() {
	hostname := "example.com"
	startPort := 1
	endPort := 100

	fmt.Printf("Scanning ports on %s...\n", hostname)

	var openPorts []int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()

			target := fmt.Sprintf("%s:%d", hostname, p)
			conn, err := net.DialTimeout("tcp", target, 500*time.Millisecond)

			if err == nil {
				conn.Close()
				mutex.Lock()
				openPorts = append(openPorts, p)
				mutex.Unlock()
			}
		}(port)
	}

	wg.Wait()
	sort.Ints(openPorts)

	fmt.Println("Open ports:")
	for _, port := range openPorts {
		fmt.Printf("%d - %s\n", port, getServiceName(port))
	}
}

func getServiceName(port int) string {
	serviceName := "Unknown"

	if p, err := parse.LookupPort(port, "tcp"); err == nil {
		serviceName = p.Name
	}

	return serviceName
}
