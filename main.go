package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	/*remote, err := url.Parse("http://localhost:81")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}*/

	//http.HandleFunc("localhost/", serveLocal)
	//http.HandleFunc("0.0.0.0/", proxy)
	//http.HandleFunc("henro.ga/", serveHenro)
	//if err := http.ListenAndServe(":81", nil); err != nil {
	//	log.Fatal(err)
	//}
	proxy()
}

func proxy() {
	remote, err := url.Parse("https://henry.pink")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	// use http.Handle instead of http.HandleFunc when your struct implements http.Handler interface
	http.Handle("/", &ProxyHandler{proxy})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("X-Ben", "Rad")
	ph.p.ServeHTTP(w, r)
}
