package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectSimple(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	dest := "http://example.com"

	redirector := &Redirector{
		DestURL: dest,
	}
	redirector.Redirect(w, req)

	if w.Code != 302 {
		t.Fatalf("expected status 302; received %d", w.Code)
	}

	if w.HeaderMap["Location"][0] != dest {
		t.Fatalf("expected location %s; received %s", dest, w.HeaderMap["Location"])
	}
}

func TestRedirectWithSubdir(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	dest := "http://example.com/foo"

	redirector := &Redirector{
		DestURL: dest,
	}
	redirector.Redirect(w, req)

	if w.Code != 302 {
		t.Fatalf("expected status 302; received %d", w.Code)
	}

	if w.HeaderMap["Location"][0] != dest {
		t.Fatalf("expected location %s; received %s", dest, w.HeaderMap["Location"])
	}
}

func TestRedirectWithPath(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/a/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	dest := "http://example.com/foo"

	redirector := &Redirector{
		DestURL: dest,
	}
	redirector.Redirect(w, req)

	if w.Code != 302 {
		t.Fatalf("expected status 302; received %d", w.Code)
	}

	tgt := dest + "/a/path"

	if w.HeaderMap["Location"][0] != tgt {
		t.Fatalf("expected location %s; received %s", tgt, w.HeaderMap["Location"])
	}
}

func TestRedirectWithSubdirAndPath(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/a/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	dest := "http://example.com"

	redirector := &Redirector{
		DestURL: dest,
	}
	redirector.Redirect(w, req)

	if w.Code != 302 {
		t.Fatalf("expected status 302; received %d", w.Code)
	}

	tgt := dest + "/a/path"

	if w.HeaderMap["Location"][0] != tgt {
		t.Fatalf("expected location %s; received %s", tgt, w.HeaderMap["Location"])
	}
}
