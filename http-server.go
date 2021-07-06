// +build ignore 
package main 

import (
    "fmt"
    "log"
    "net/http"
)

// http.Request represents the client HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
    // r.URL.Path is the path component of the request URL.
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

