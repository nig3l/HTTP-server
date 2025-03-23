package main

import (
	"fmt"
	"log"
	"net/http"
    "time"
    
)

func logRequest(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        next(w, r)
        log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(startTime))
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "Welcome to the home page!")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}