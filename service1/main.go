package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var alive = true

func handler(w http.ResponseWriter, r *http.Request) {
	if !alive {
		panic("crash")
	}
	fmt.Fprintln(os.Stdout, "URL:", r.URL.String())

	name := os.Getenv("SERVICE_NAME")
	// Ignore errs in demo
	host, _ := os.Hostname()
	addrs, _ := net.LookupHost(host)

	fmt.Fprintf(w, "Hello from %s\n", name)
	fmt.Fprintf(w, "host: %s\n", host)
	fmt.Fprintf(w, "addrs: %s\n", addrs)
}

func sleeper(w http.ResponseWriter, r *http.Request) {
	alive = false
	go func() {
		time.Sleep(35 * time.Second)
		alive = true
	}()
}

func main() {
	http.HandleFunc("/service/1/sleep", sleeper)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
