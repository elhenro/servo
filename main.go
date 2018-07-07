package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func dom1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey dude!")
}

func dom2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We welcome you, sire!")
}

func main() {
	r := mux.NewRouter()
	d1 := r.Host("adomain.com").Subrouter()
	d2 := r.Host("seconddomain.com").Subrouter()

	d1.HandleFunc("/", dom1Handler)
	d2.HandleFunc("/", dom2Handler)

	http.ListenAndServe(":8080", nil)
}
