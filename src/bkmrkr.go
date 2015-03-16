package main

import (
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
    "time"
)

var urls []string = make([]string, 1)

type Page struct {
    Title string
    Urls []string
}

func index(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    t, _ := template.ParseFiles("templates/index.html")

    t.Execute(w, &Page{Title: "bkmrkr", Urls: urls})
}

func add(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)

    url := r.FormValue("url")
    urls = append(urls, url)
    log.Println("Add:  '" + url + "'")

    http.Redirect(w, r, "/index", 302)
}

func main() {
    requestRouter := mux.NewRouter()
    requestRouter.HandleFunc("/", index).Methods("GET")
    requestRouter.HandleFunc("/index", index).Methods("GET")
    requestRouter.HandleFunc("/add", add).Methods("POST")
    // static content
    requestRouter.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img/"))))
    requestRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css/"))))
    requestRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js/"))))

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
