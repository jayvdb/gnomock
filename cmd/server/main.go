package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/orlangure/gnomock/internal/gnomockd"
)

var version string

func main() {
	var (
		v       bool
		verbose bool
		port    int
	)

	flag.BoolVar(&v, "v", false, "display current version")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")
	flag.IntVar(&port, "port", 23042, "gnomockd port number")
	flag.Parse()

	if v {
		fmt.Println(version)
		os.Exit(0)
	}

	if verbose {
		log.Println("Verbose logging")
	}

	if pStr, ok := os.LookupEnv("GNOMOCKD_PORT"); ok {
		if p, err := strconv.Atoi(pStr); err == nil {
			port = p
		}
	}

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	log.Println("Binding to ", addr)
	handler := gnomockd.Handler()
	log.Println("Created handler")
	rv := http.ListenAndServe(addr, handler)
	log.Fatal("listener returned", rv)
}
