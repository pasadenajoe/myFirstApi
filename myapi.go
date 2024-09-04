package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pasadenajoe/myFirstApi/rutas"
)

// --------------------------------------------------
// * * * * * * * * * *
// --------------------------------------------------
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rutas.HomeHandler)
	r.HandleFunc("/est/getById/{id}", rutas.EstGetByIdHandler).Methods("GET")
	r.HandleFunc("/est/getByCed/{ced}", rutas.EstGetByCedHandler).Methods("GET")
	r.HandleFunc("/est/insert", rutas.EstInsertHandler).Methods("POST")
	r.HandleFunc("/est/update", rutas.EstUpdateHandler).Methods("PUT")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8095",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
