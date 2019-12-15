package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/valentergs/books_backend/controllers"
	"github.com/valentergs/books_backend/driver"
)

var db *sql.DB

func main() {
	db := driver.ConnectDB()
	livroctl := controllers.ControllerLivro{}

	// gorilla.mux
	router := mux.NewRouter()

	// LIVRO URL ====================================
	router.HandleFunc("/", livroctl.TodosLivros(db)).Methods("GET")
	router.HandleFunc("/{id}", livroctl.LivroUnico(db)).Methods("GET")
	router.HandleFunc("/deletar/{id}", livroctl.LivroApagar(db)).Methods("DELETE")
	router.HandleFunc("/editar/{id}", livroctl.LivroEditar(db)).Methods("PUT")

	log.Println("Listen on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
