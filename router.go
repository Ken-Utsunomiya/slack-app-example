package main

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/slack/command", NewHandler().ServeHTTP)
	return mux
}
