package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Go Language</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("server is running...")
	http.ListenAndServe(":8080", nil)
}
