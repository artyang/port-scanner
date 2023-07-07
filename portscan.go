package main

import (
    "fmt"
    "net"
    "time"
)

func scanPort(protocol, hostname string, port int) string {
    address := fmt.Sprintf("%s:%d", hostname, port)
    conn, err := net.DialTimeout(protocol, address, 1*time.Second)
    if err != nil {
        return "Closed"
    }
    defer conn.Close()
    return "Open"
}

func main() {
    hostname := "localhost"
    protocols := []string{"tcp", "udp"}
    ports := []int{80, 443, 8080}

    for _, protocol := range protocols {
        for _, port := range ports {
            result := scanPort(protocol, hostname, port)
            fmt.Printf("Port %d/%s: %s\n", port, protocol, result)
        }
    }
}
