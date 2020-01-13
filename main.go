package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

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
	router.HandleFunc("/inserir", livroctl.LivroInserir(db)).Methods("POST")
	router.HandleFunc("/deletar/{id}", livroctl.LivroApagar(db)).Methods("DELETE")
	router.HandleFunc("/editar/{id}", livroctl.LivroEditar(db)).Methods("PUT")

	// CORS ==========================================================
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	// Insert the middleware
	handler := c.Handler(router)

	log.Println("Listen on port 8080...")
	// Quando  usar o CORS colocar "handler" no lugar do "router"
	// Quando usar ROUTER usar "router"
	log.Fatal(http.ListenAndServe(":8080", handler))
}
