package models

import "github.com/valbercarreiro/alura/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectAllProducts, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()

	deletarProduto, err := db.Prepare("delete from produtos where id =  $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(idProduto)
	defer db.Close()
}

func ConsultarProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()

	produtoRetorno, err := db.Query("select * from produtos where id=$1", idProduto)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for produtoRetorno.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoRetorno.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Preco = preco
	}
	defer db.Close()
	return produto
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
