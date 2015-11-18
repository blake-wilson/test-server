package main

import (
	"fmt"
	"net/http"
)

// Simple web server
var counter int
var client *http.Client

func init() {
	client = &http.Client{}
}

func failHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	w.WriteHeader(http.StatusOK)
}

func getter(addresses ...string) error {
	// use client to make HTTP request for each address
	for _, addr := range addresses {
		resp, err := client.Get(addr)
		if err != nil || resp.StatusCode != http.StatusOK {
			return fmt.Errorf("HTTP call not OK")
		}
	}
	return nil
}

func main() {
	http.HandleFunc("/", successHandler)
	http.ListenAndServe(":8080", nil)
}
