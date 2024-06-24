package repositories

import (
    "database/sql"
    "errors"
    "log"
    "pedido-gerenciamento/models"
    _ "github.com/lib/pq"
)

func SalvarProduto(produto models.Produto) error {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return err
    }
    defer db.Close()

    query := `INSERT INTO produtos (nome, quantidade, preco) VALUES ($1, $2, $3)`
    _, err = db.Exec(query, produto.Nome, produto.Quantidade, produto.Preco)
    if err != nil {
        return err
    }
    return nil
}

func ListarProdutos() []models.Produto {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, nome, quantidade, preco FROM produtos")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var produtos []models.Produto
    for rows.Next() {
        var produto models.Produto
        if err := rows.Scan(&produto.ID, &produto.Nome, &produto.Quantidade, &produto.Preco); err != nil {
            log.Fatal(err)
        }
        produtos = append(produtos, produto)
    }
    return produtos
}

func BuscarProduto(nome string) (models.Produto, error) {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return models.Produto{}, err
    }
    defer db.Close()

    var produto models.Produto
    query := `SELECT id, nome, quantidade, preco FROM produtos WHERE nome=$1`
    err = db.QueryRow(query, nome).Scan(&produto.ID, &produto.Nome, &produto.Quantidade, &produto.Preco)
    if err != nil {
        if err == sql.ErrNoRows {
            return produto, errors.New("produto não encontrado")
        }
        return produto, err
    }
    return produto, nil
}

func BuscarProdutoPorID(id int) (models.Produto, error) {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return models.Produto{}, err
    }
    defer db.Close()

    var produto models.Produto
    query := `SELECT id, nome, quantidade, preco FROM produtos WHERE id=$1`
    err = db.QueryRow(query, id).Scan(&produto.ID, &produto.Nome, &produto.Quantidade, &produto.Preco)
    if err != nil {
        if err == sql.ErrNoRows {
            return produto, errors.New("produto não encontrado")
        }
        return produto, err
    }
    return produto, nil
}

func AtualizarEstoque(produtoNome string, quantidade int) error {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return err
    }
    defer db.Close()

    query := `UPDATE produtos SET quantidade = quantidade + $1 WHERE nome = $2`
    _, err = db.Exec(query, quantidade, produtoNome)
    if err != nil {
        return err
    }
    return nil
}

func BuscarProdutoPorNome(nome string) (models.Produto, error) {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return models.Produto{}, err
    }
    defer db.Close()

    var produto models.Produto
    query := `SELECT id, nome, quantidade, preco FROM produtos WHERE nome=$1`
    err = db.QueryRow(query, nome).Scan(&produto.ID, &produto.Nome, &produto.Quantidade, &produto.Preco)
    if err != nil {
        if err == sql.ErrNoRows {
            return produto, errors.New("produto não encontrado")
        }
        return produto, err
    }
    return produto, nil
}
