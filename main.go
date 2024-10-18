package main

import (
    "fmt"
    "log"
    "net/http"
)

// Serve static HTML for frontend
func serveFrontend(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// Handle textarea input submission
func handleText(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        text := r.FormValue("inputText")
        // Process the text here (for example, highlight non-breaking spaces)
        fmt.Fprintf(w, "Processed text: %s", text)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/", serveFrontend)
    http.HandleFunc("/process", handleText)

    port := ":8080"
    log.Printf("Starting server on %s...", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
