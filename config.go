package main

import "net/http"

func Config(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Accept", "text/json")
		w.Header().Set("Accept-Charset", "utf-8")
		w.Header().Set("Accept-Encoding", "gzip, deflate")
		inner.ServeHTTP(w, r)
	})
}
