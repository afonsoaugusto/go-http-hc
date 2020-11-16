package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World! %s", time.Now())
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
	fmt.Println("500")
	fmt.Fprintf(w, "500")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/health", health)
	http.HandleFunc("/health-check", health)
	http.HandleFunc("/healthcheck", health)
	fmt.Println("Listening 8080")
	http.ListenAndServe(":8080", nil)
}
