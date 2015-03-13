package main

import (
    "html/template"
    "log"
    "net/http"
    "time"
)

type Page struct {
    Title string
}

func index(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    t, _ := template.ParseFiles("templates/index.html")
    t.Execute(w, &Page{Title: "bkmrkr"})
}

func status(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    t, _ := template.ParseFiles("templates/status.html")
    t.Execute(w, &Page{Title: "Status"})
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
    mux.HandleFunc("/status", status)

    server := &http.Server{
        Addr:           ":8000",
        Handler:        mux,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    log.Println("Listening...")
    log.Fatal(server.ListenAndServe())
}
