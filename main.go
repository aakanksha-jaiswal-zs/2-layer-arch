package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	handlers "sample-api/handlers/customer"
	stores "sample-api/stores/customer"
)

var db = map[string]int{
	"Aryan":  1,
	"Umang":  2,
	"Ruchit": 3,
}

func main() {
	store := stores.New(db)
	h := handlers.New(store)

	r := mux.NewRouter()

	r.HandleFunc("/hello", h.Get).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
