package main

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "bkmrkr")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    http.ListenAndServe(":8000", mux)
}
