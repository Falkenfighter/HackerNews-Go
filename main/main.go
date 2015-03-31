package main
import (
    "net/http"
    "log"
    "./lib"
)

func main() {
    // Init routes
    http.HandleFunc("/", lib.Index)
    http.HandleFunc("/toprated", lib.TopRated)

    // Init server
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
