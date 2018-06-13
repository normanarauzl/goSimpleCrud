package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/normanarauzl/goSimpleCrud/templates/views/Employee"
)

func main() {
	log.Println("Server Iniciado en http://localhost:3000")
	mux := http.NewServeMux()
	mux.HandleFunc("/", Employee.GoIndex)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
