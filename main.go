package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	cookieName = "pressMe"
)

func pressme(w http.ResponseWriter, _ *http.Request) {
	formCookie := &http.Cookie{
		Name:     cookieName,
		Value:    time.Now().Format("15:04:05"),
		Path:     "/",
		Expires:  time.Now().Add(5 * time.Minute),
		MaxAge:   300,                     // 5 минут
		Secure:   false,                   // yet 'false' as TLS is not used
		HttpOnly: true,                    // 'true' secures from XSS attacks
		SameSite: http.SameSiteStrictMode, // base CSRF attack protection
	}

	http.SetCookie(w, formCookie)

	_, err := fmt.Fprintf(w, "<html><body><a href=/pressme>Press me</a><body></html>")
	if err != nil {
	    log.Println(err)
	}
}

func cookie(w http.ResponseWriter, r *http.Request) {
	var err error
	var pressMeCookie *http.Cookie

	pressMeCookie, err = r.Cookie(cookieName)
	if err != nil {
		return
	}

	fmt.Println("A cookie Detected")
	cookieValue := pressMeCookie.Value
	_, err = fmt.Fprintf(w, "You pressed me on %s!", cookieValue)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", pressme)
	http.HandleFunc("/pressme", cookie)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
