package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/", test)
	//http.HandleFunc("http://192.168.2.135:8050", test)

	log.Fatal(http.ListenAndServe(":8050", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}
func test(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "test")
	remote, err := url.Parse("http://localhost:8047")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.Handle("/test/test", &ProxyHandler{proxy})
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("X-Ben", "Rad")
	ph.p.ServeHTTP(w, r)
}
