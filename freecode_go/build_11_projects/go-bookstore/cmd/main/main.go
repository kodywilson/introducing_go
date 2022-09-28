package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kodywilson/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

// Testing with Curl - first install mysql if needed, create user and grant privileges
// Create a book
// curl -d '{"name":"Go Is Fun","author":"Nawt Reel", "publication":"Fresh Press"}' -H 'Content-Type: application/json' http://localhost:9010/book/
// Get book by id
// curl -v http://localhost:9010/book/3
