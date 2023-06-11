package main

import (
	"log"
	"net/http"

	"github.com/albar2305/go-book/config"
	"github.com/albar2305/go-book/controllers/bookcontroller"
	"github.com/albar2305/go-book/controllers/homecontroler"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroler.Welcome)

	http.HandleFunc("/books", bookcontroller.Index)
	http.HandleFunc("/books/add", bookcontroller.Add)
	http.HandleFunc("/books/edit", bookcontroller.Edit)
	http.HandleFunc("/books/detail", bookcontroller.Detail)
	http.HandleFunc("/books/delete", bookcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
