package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	return
}
