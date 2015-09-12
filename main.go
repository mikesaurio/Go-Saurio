package main

import (
    "fmt"
    "net/http"
)

type Hello struct{
NAME string
}

func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "¡Hola!" , h.NAME)
}

func main() {
    var h Hello
    h.NAME = "Mike"
    http.ListenAndServe("localhost:4000", h)
}

