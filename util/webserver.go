package util

import (
	"log"
	"net/http"
	"strings"
	"time"
)

const httpsPort = ":8443"
const httpPort = ":8080"

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
	log.Fatal(http.ListenAndServeTLS(httpsPort, "cert.pem", "key.pem", nil))
}

func startHTTPRedirectServer() {
	log.Fatal(http.ListenAndServe(httpPort, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := "https://" + strings.Split(r.Host, ":")[0] + httpsPort + r.URL.Path
		if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, target, http.StatusMovedPermanently)
	})))
}
