package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"path"
)

var (
	destUrl string
	code    int
)

type (
	Redirector struct {
		DestURL string
	}
)

func init() {
	flag.StringVar(&destUrl, "url", "", "URL for redirect")
	flag.IntVar(&code, "code", 302, "Status code (default: 302)")
	flag.Parse()
}

func (r *Redirector) Redirect(w http.ResponseWriter, req *http.Request) {
	log.Printf("redirecting request from client %s", req.RemoteAddr)
	u, err := url.Parse(r.DestURL)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dUrl := &url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   path.Join(u.Path, req.URL.Path),
	}
	http.Redirect(w, req, dUrl.String(), code)
}

func main() {
	if destUrl == "" {
		log.Fatal("you must specify a URL")
	}

	redirector := &Redirector{
		DestURL: destUrl,
	}

	http.HandleFunc("/", redirector.Redirect)

	log.Printf("Listening on 8080; redirecting to %s", destUrl)
	http.ListenAndServe(":8080", nil)
}
