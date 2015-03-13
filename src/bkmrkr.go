package main

import (
    "github.com/gorilla/mux"
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

func add(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    log.Println("TODO: add bookmark")
    index(w, r)
}


func main() {
    requestRouter := mux.NewRouter()
    requestRouter.HandleFunc("/", index).Methods("GET")
    requestRouter.HandleFunc("/status", status).Methods("GET")
    requestRouter.HandleFunc("/add", add).Methods("POST")

    server := &http.Server{
        Addr:           ":8000",
        Handler:        requestRouter,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    log.Println("Listening...")
    log.Fatal(server.ListenAndServe())
}
