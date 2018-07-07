package main

import (
	"fmt"
	"net/http"
)

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404 - whoa.. dude! page not found.")
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoa.. dude! hi there")
}
func poweruserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoa.. dude! hi there")
}

func main() {
	http.HandleFunc("/", notfoundHandler)
	http.HandleFunc("/home", homeHandler)

	// from different domain
	http.HandleFunc("0.0.0.0", poweruserHandler)

	http.ListenAndServe(":8080", nil)
}
