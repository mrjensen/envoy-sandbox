package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("SERVICE_NAME")
	// Ignore errs in demo
	host, _ := os.Hostname()
	addrs, _ := net.LookupHost(host)

	fmt.Fprintf(w, "Hello from %s\n", name)
	fmt.Fprintf(w, "host: %s\n", host)
	fmt.Fprintf(w, "addrs: %s\n", addrs)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
