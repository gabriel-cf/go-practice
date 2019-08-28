package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, s string) {
	fmt.Fprintf(w, "Hello, " + s)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Yo")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}