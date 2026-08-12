package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "hello %s", name)
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			name = "stranger"
			q := r.URL.Query()
			q.Set("name", name)
			r.URL.RawQuery = q.Encode()
		}
		next(w, r)
	}
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if strings.ContainsAny(name, "0123456789йцукенгшщзхъфывапролджэячсмитьбюё ") {
			name = "dirty hacker"
			q := r.URL.Query()
			q.Set("name", name)
			r.URL.RawQuery = q.Encode()
		}
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/hello", SetDefaultName(Sanitize(HelloHandler)))

	http.ListenAndServe(":8080", nil)
}
