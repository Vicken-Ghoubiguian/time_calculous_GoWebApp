package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
  hostname, err := os.Hostname()
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, world!\nVersion: 1.0.0\n" + hostname + "\n")
  })

  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":8080", nil)
}
