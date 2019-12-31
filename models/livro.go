package models

//Livro is an exportable type
type Livro struct {
	ID             int    `json:"livro_id"`
	Titulo         string `json:"titulo"`
	TituloOriginal string `json:"titulo_original"`
	Autor          string `json:"autor"`
	Tradutor       string `json:"tradutor"`
	Isbn           string `json:"isbn"`
	Cdd            string `json:"cdd"`
	Cdu            string `json:"cdu"`
	Ano            string `json:"ano"`
	Tema           string `json:"tema"`
	Editora        string `json:"editora"`
	Paginas        string `json:"paginas"`
	Idioma         string `json:"idioma"`
	Formato        string `json:"formato"`
	Dono           string `json:"dono"`
	Photourl       string `json:"photourl"`
}
