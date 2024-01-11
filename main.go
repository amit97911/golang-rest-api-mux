package main

import (
	"books"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to Books management. Server is running on port 8000")
	})

	books.ProductsModule(router)

	defer log.Fatal(srv.ListenAndServe())
}
