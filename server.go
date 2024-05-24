package main

import (
    "crypto/sha1"
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Fprintf(w, "<html><body>Error - unable to get hostname: %v</body></html>", err)
        return
    }

    hash := sha1.New()
    hash.Write([]byte(hostname))
    hashBytes := hash.Sum(nil)
    color := fmt.Sprintf("#%02x%02x%02x", hashBytes[0], hashBytes[1], hashBytes[2]) 

    fmt.Fprintf(w, "<html><body style='background-color:%s'>Container ID: %s</body></html>", color, hostname)
    fmt.Fprintf(w, "<html><body>Testing of Jenkins (starting by new push on GitHub)</body></html>")
}


func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":1234", nil)
}
