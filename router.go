package main

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/slash", NewHandler().ServeHTTP)
	return mux
}
