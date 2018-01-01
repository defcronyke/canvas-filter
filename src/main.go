package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func init() {
  r := mux.NewRouter()
  r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/dist")))).Methods("GET")
  http.Handle("/", r)
}
