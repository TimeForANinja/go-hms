package util

import (
	"log"
	"net/http"
	"time"
)

type page struct {
	body       []byte
	statusCode int
}

func sendPage(w http.ResponseWriter, page *page) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "yey", Value: "Boy", Secure: true, HttpOnly: true, Expires: expiration, MaxAge: 24 * 60 * 60}
	http.SetCookie(w, &cookie)

	w.WriteHeader(page.statusCode)
	w.Write(page.body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	sendPage(w, &page{[]byte("Hello World"), 200})
}

// StartWeb starts the web part of the home media server
func StartWeb() {
	go startHTTPRedirectServer()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}

func startHTTPRedirectServer() {
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := "https://" + r.Host + r.URL.Path
		if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, target, http.StatusMovedPermanently)
	})))
}
