package main

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(host string, port int) bool {

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		time.Second,
	)

	if err != nil {
		return false
	}

	conn.Close()

	return true
}
