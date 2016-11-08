package main

import (
	"encoding/json"
	"net/http"
)

func ResponseHandler(context interface{}, status int, w http.ResponseWriter) {
	js, err := json.Marshal(context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(js)
}
