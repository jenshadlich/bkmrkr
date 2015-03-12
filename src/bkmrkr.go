package main

import (
    "html/template"
    "log"
    "net/http"
)

type Page struct {
    Title string
}

func index(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Header)
    t, _ := template.ParseFiles("templates/index.html")
    t.Execute(w, &Page{Title: "bkmrkr"})
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)

    log.Println("Listening...")
    http.ListenAndServe(":8000", mux)
}
