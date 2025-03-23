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

    html := `
    <!DOCTYPE html>
    <html>
        <head>
            <title>Cool HTTP Server</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    max-width: 800px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: #f0f2f5;
                }
                .container {
                    background-color: white;
                    padding: 20px;
                    border-radius: 8px;
                    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
                }
                .button {
                    background-color: #4CAF50;
                    color: white;
                    padding: 10px 20px;
                    border: none;
                    border-radius: 4px;
                    cursor: pointer;
                }
                .button:hover {
                    background-color: #45a049;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Welcome to Our Cool HTTP Server!</h1>
                <p>This is a simple but modern-looking interface for our server.</p>
                <button class="button" onclick="window.location.href='/path/test'">Try Path Handler</button>
            </div>
        </body>
    </html>
    `
    fmt.Fprintf(w, html)
    
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
    html := fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
        <head>
            <title>Path Info</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    max-width: 800px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: #f0f2f5;
                }
                .container {
                    background-color: white;
                    padding: 20px;
                    border-radius: 8px;
                    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
                }
                .path-info {
                    background-color: #e9ecef;
                    padding: 10px;
                    border-radius: 4px;
                    font-family: monospace;
                }
                .back-button {
                    background-color: #6c757d;
                    color: white;
                    padding: 10px 20px;
                    border: none;
                    border-radius: 4px;
                    cursor: pointer;
                    text-decoration: none;
                    display: inline-block;
                    margin-top: 20px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Path Information</h1>
                <div class="path-info">
                    You've requested: %s
                </div>
                <a href="/" class="back-button">Back to Home</a>
            </div>
        </body>
    </html>
    `, r.URL.Path)
    fmt.Fprintf(w, html)
    
}

func main() {

    http.HandleFunc("/", logRequest(homeHandler))
    http.HandleFunc("/path/", logRequest(pathHandler))
    
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}