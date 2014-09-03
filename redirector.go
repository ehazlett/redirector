package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	URL  string
	CODE int
)

func init() {
	flag.StringVar(&URL, "url", "", "URL for redirect")
	flag.IntVar(&CODE, "code", 302, "Status code (default: 302)")
	flag.Parse()
}

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("redirecting request from client %s", r.RemoteAddr)
	http.Redirect(w, r, URL, CODE)
}

func main() {
	if URL == "" {
		log.Fatal("you must specify a URL")
	}
	http.HandleFunc("/", redirect)
	log.Printf("Listening on 8080; redirecting to %s", URL)
	http.ListenAndServe(":8080", nil)
}
