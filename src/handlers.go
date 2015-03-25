package main

import (
    "html/template"
    "log"
    "net/http"
)

var urls []string = make([]string, 1)

type Page struct {
    Title string
    Urls []string
}

func Index(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    t, _ := template.ParseFiles("templates/index.html")

    t.Execute(w, &Page{Title: "bkmrkr", Urls: urls})
}

func Add(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)

    url := r.FormValue("url")
    urls = append(urls, url)
    log.Println("Add:  '" + url + "'")

    http.Redirect(w, r, "/index", 302)
}

