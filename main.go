package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Go Language!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Tp get in touch, please send an email to <a "+
		"href=\"mohammadreza.che.eng@gmail.com\">support@mohammadreza.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println("This is a FAQ Page")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find what you are looking for!</h1>")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	// 404 page
	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":8080", r)
}
