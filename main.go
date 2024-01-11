package main

import (
	"books"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	err    error
	router *mux.Router
	srv    *http.Server
)

func main() {
	db, err = gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// db.AutoMigrate(&books.Books{})

	router = mux.NewRouter().StrictSlash(true)

	srv = &http.Server{
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

	db.Migrator().CreateTable(&books.Books{})

	defer log.Fatal(srv.ListenAndServe())
}
