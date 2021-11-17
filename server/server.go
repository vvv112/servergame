package main

import (
	"fmt"
	"net/http"
	)

func main() {
	println("server 0.0.1")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "server1")
		})
		http.ListenAndServe(":80", nil)

}