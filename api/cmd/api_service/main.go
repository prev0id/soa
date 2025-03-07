package main

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	usersProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "user-service:8081"

			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/user")
			if req.URL.Path == "" {
				req.URL.Path = "/"
			}

			req.Host = req.URL.Host
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/user/") {
			usersProxy.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
