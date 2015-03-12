package main

import (
    "io"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "bkmrkr")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
    http.ListenAndServe(":8000", mux)
}
