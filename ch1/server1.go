package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("172.16.10.30:8123", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "Request from %q \n", r.RemoteAddr)
	fmt.Fprintf(w, "URL.Path = %q", r.URL)
}
