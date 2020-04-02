package controllers

import (
	"net/http"
	"html/template"
	"github.com/brendonhps/models"
	"strconv"
	"log"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index : 
func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

// Store : 
func Store(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

// Insert : 
func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preco: ",  err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao da quantidade: ",  err)
		}

		models.CriaNovoProduto(nome,descricao,precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

// Delete :
func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r , "/", 301)
}

//Edit :
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	produto := models.EditaProduto(idDoProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

//Update :
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome:= r.FormValue("nome")
		descricao:= r.FormValue("descricao")
		preco:= r.FormValue("preco")
		quantidade:= r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)

		if err !=  nil {
			log.Println("Erro na conversao do Id")
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err !=  nil {
			log.Println("Erro na conversao do preco")
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err !=  nil {
			log.Println("Erro na conversao do quantidade")
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/", 301)

}	