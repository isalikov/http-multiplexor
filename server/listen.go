package server

import (
    "net/http"
    "os"
)

func Listen() {
    addr := os.Getenv("ADDR")

    if addr == "" {
        addr = ":8080"
    }

    http.HandleFunc("/", HandleFunc)

    http.ListenAndServe(addr, nil)
}
