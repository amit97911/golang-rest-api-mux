module github.com/amit97911/golang-rest-api-mux

go 1.21.5

require github.com/gorilla/mux v1.8.1

replace books v1.0.0 => ./modules/books

require books v1.0.0
