package main

import (
    "fmt"
    "net/http"
)

func health(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "OK")
}

func v1(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "WSC")
}

func main() {
    http.HandleFunc("/health", health)
    http.HandleFunc("/v1", v1)
    http.ListenAndServe(":8000", nil)
}