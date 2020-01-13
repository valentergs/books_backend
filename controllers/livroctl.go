package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/valentergs/books_backend/models"
	"github.com/valentergs/books_backend/utils"
	// "github.com/gocolly/colly"
)

//ControllerLivro será exportado
type ControllerLivro struct{}

//TodosLivros será exportado ==========================================
func (c ControllerLivro) TodosLivros(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error

		if r.Method != "GET" {
			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		rows, err := db.Query("SELECT * FROM livros ORDER BY livro_id DESC;")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			fmt.Println(err)
			return
		}

		defer rows.Close()

		clts := make([]models.Livro, 0)
		for rows.Next() {
			clt := models.Livro{}
			err := rows.Scan(&clt.ID, &clt.Titulo, &clt.TituloOriginal, &clt.Autor, &clt.Tradutor, &clt.Isbn, &clt.Cdd, &clt.Cdu, &clt.Ano, &clt.Tema, &clt.Editora, &clt.Paginas, &clt.Idioma, &clt.Formato, &clt.Dono, &clt.Photourl)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			clts = append(clts, clt)
			
		}
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, clts)

	}

}

//LivroUnico será exportado ==========================================
func (c ControllerLivro) LivroUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var livro models.Livro

		if r.Method != "GET" {
			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			erro.Message = "Numero ID inválido"
		}

		row := db.QueryRow("SELECT * FROM livros WHERE livro_id=$1;", id)

		err = row.Scan(&livro.ID, &livro.Titulo, &livro.TituloOriginal, &livro.Autor, &livro.Tradutor, &livro.Isbn, &livro.Cdd, &livro.Cdu, &livro.Ano, &livro.Tema, &livro.Editora, &livro.Paginas, &livro.Idioma, &livro.Formato, &livro.Dono, &livro.Photourl)
		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, livro)

	}

}

//LivroInserir será exportado ===========================================
func (c ControllerLivro) LivroInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var livro models.Livro
		// var link string

		if r.Method != "POST" {
			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		json.NewDecoder(r.Body).Decode(&livro)
		
		isbn := livro.Isbn
		link := utils.PhotoLink(isbn)

		expressaoSQL := `INSERT INTO livros (titulo, titulo_original, autor, tradutor, isbn, cdd, cdu, ano, tema, editora, paginas, idioma, formato, dono, photourl) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);`
		_, err := db.Exec(expressaoSQL, livro.Titulo, livro.TituloOriginal, livro.Autor, livro.Tradutor, livro.Isbn, livro.Cdd, livro.Cdu, livro.Ano, livro.Tema, livro.Editora, livro.Paginas, livro.Idioma, livro.Formato, livro.Dono, link)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM livros WHERE isbn=$1;", livro.Isbn)
		err = row.Scan(&livro.ID, &livro.Titulo, &livro.TituloOriginal, &livro.Autor, &livro.Tradutor, &livro.Isbn, &livro.Cdd, &livro.Cdu, &livro.Ano, &livro.Tema, &livro.Editora, &livro.Paginas, &livro.Idioma, &livro.Formato, &livro.Dono, &livro.Photourl)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, livro)

	}
}

//LivroApagar será exportado =========================================
func (c ControllerLivro) LivroApagar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "DELETE" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo usuario no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		db.QueryRow("DELETE FROM livros where usuario_id=$1;", id)

		SuccessMessage := "Livro deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		utils.ResponseJSON(w, SuccessMessage)

	}
}

//LivroEditar será exportado =========================================
func (c ControllerLivro) LivroEditar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var livro models.Livro
		var error models.Error

		if r.Method != "PUT" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		json.NewDecoder(r.Body).Decode(&livro)

		expressaoSQL := `UPDATE livros SET titulo=$1, titulo_original=$2, autor=$3, tradutor=$4, isbn=$5, cdd=$6, cdu=$7, ano=$8, tema=$9, editora=$10, paginas=$11, idioma=$12, formato=$13, dono=$14, photourl=$15 WHERE livro_id=$16;`
		_, err = db.Exec(expressaoSQL, livro.Titulo, livro.TituloOriginal, livro.Autor, livro.Tradutor, livro.Isbn, livro.Cdd, livro.Cdu, livro.Ano, livro.Tema, livro.Editora, livro.Paginas, livro.Idioma, livro.Formato, livro.Dono, livro.Photourl, id)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM livros WHERE livro_id=$1;", livro.ID)
		err = row.Scan(&livro.ID, &livro.Titulo, &livro.TituloOriginal, &livro.Autor, &livro.Tradutor, &livro.Isbn, &livro.Cdd, &livro.Cdu, &livro.Ano, &livro.Tema, &livro.Editora, &livro.Paginas, &livro.Idioma, &livro.Formato, &livro.Dono, &livro.Photourl)

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, livro)

	}
}
