package models

import (
	"github.com/brendonhps/db"
)
// Produto : data struct to save the product 
type Produto struct { 
	ID int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

// BuscaTodosOsProdutos : 
func BuscaTodosOsProdutos() []Produto{

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()


	p := Produto{}
	produtos := []Produto{}
	 
	for selectDeTodosOsProdutos.Next(){
		var id , quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id , &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.ID = id
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	return produtos
}

// CriaNovoProduto :
func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome,descricao,preco,quantidade)
}

// DeletaProduto :
func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	defer db.Close()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")


	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
}


// EditaProduto :
func EditaProduto(id string) Produto {
	db:= db.ConectaComBancoDeDados()

	defer db.Close()

	produtoDoBanco , err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}
	
	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.ID = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	return produtoParaAtualizar

}

// AtualizaProduto :
func AtualizaProduto (id int, nome string, descricao string, preco float64,quantidade int) {
	db:= db.ConectaComBancoDeDados()

	defer db.Close()

	AtualizaProd, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	} 

	AtualizaProd.Exec(nome,descricao,preco,quantidade,id)
}