module github.com/amit97911/golang-rest-api-mux

go 1.21.6

require github.com/gorilla/mux v1.8.1

replace initializer v1.0.0 => ./initializer

replace books v1.0.0 => ./modules/books

require (
	books v1.0.0
	initializer v1.0.0
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	gorm.io/driver/sqlite v1.5.4 // indirect
	gorm.io/gorm v1.25.5 // indirect
)
