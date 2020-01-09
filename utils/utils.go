package utils

import (
	"encoding/json"
	"net/http"

	"github.com/valentergs/books_backend/models"
	"github.com/gocolly/colly"
)

//ResponseJSON will be exported ========================================
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//RespondWithError will be exported ====================================
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
	return
}

//PhotoLink will be exported ====================================
func PhotoLink(isbn string) string {
	c := colly.NewCollector()
	var nulink string
	c.OnHTML(".s-image", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		link = link[36:47]
		nulink = `https://images-na.ssl-images-amazon.com/images/I/` + link + `.jpg`
	})
	c.Visit(`https://www.amazon.com.br/s?k=` + isbn) 
	c.Wait()
	return nulink
}




