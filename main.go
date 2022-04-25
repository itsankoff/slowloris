package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

// cmd arguments
var flags *flag.FlagSet

// initializing flags
func init() {
	flags = flag.NewFlagSet("default", flag.ContinueOnError)
}

const legalDisclaimer = `
LEGAL DISCLAIMER: Usage of this program for attacking targets without
prior mutual consent is illegal. It is the end user's responsibility to obey
all applicable local, state and federal laws in all countries.
Developers assume no liability and are not responsible for any misuse or
damage caused by this program.
`

func printHelp() {
	flags.PrintDefaults()
	fmt.Println(legalDisclaimer)
}

func checkHelp(args []string) {
	for _, arg := range args {
		if arg == "help" || arg == "-help" || arg == "--help" {
			printHelp()
			os.Exit(0)
		}
	}
}

func main() {
	// cmd flags
	var rawURL *string = flags.String("url", "", "URL to perform attack. Format: http[s]://<domain>[:<port>]/<path>?<query-string>")
	var count *int64 = flags.Int64("count", int64(10), "Number of parallel workers")
	var interval *time.Duration = flags.Duration("interval", 1*time.Second, "Interval for sending data")
	var timeout *time.Duration = flags.Duration("timeout", 10*time.Second, "Timeout for the whole operation")
	var userAgent *string = flags.String("user-agent", "random", "Custom User-Agent header. Default 'random' which sends random header for each worker")

	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			printHelp()
		}
	}()

	flags.Parse(os.Args[1:])
	if !flags.Parsed() {
		return
	}
	checkHelp(os.Args)

	if !strings.Contains(*rawURL, "http") {
		err = errors.New("no scheme provided. use (http or https)")
		return
	}
	parsed, err := url.Parse(*rawURL)
	if err != nil {
		return
	}

	options := Options{
		URL:       parsed,
		UserAgent: *userAgent,
		Secure:    parsed.Scheme == "https",
		Count:     *count,
		Interval:  *interval,
		Timeout:   *timeout,
	}

	fmt.Println("slowloris...")
	fmt.Printf("creating zoo of %d loris for %s\n", *count, *rawURL)
	if err := Zoo(options); err != nil {
		return
	}

	return
}
