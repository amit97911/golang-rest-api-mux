package books

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductsModule(router *mux.Router) {
	book_router := router.PathPrefix("/products").Subrouter()

	book_router.HandleFunc("/", getBooks).Methods("GET")
	book_router.HandleFunc("/{id}", getBook).Methods("GET")
	book_router.HandleFunc("/", createBook).Methods("POST")
	book_router.HandleFunc("/{id}", updateBook).Methods("PUT")
	book_router.HandleFunc("/{id}", deleteBook).Methods("DELETE")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Get Books!")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Get Book by ID!")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Create a new book!")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Update a book by ID")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Delete a book by ID")
}
