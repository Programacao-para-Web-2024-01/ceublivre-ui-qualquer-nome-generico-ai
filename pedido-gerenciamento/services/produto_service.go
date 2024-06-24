package services

import (
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/repositories"
)

func AdicionarProduto(produto models.Produto) (models.Produto, error) {
    err := repositories.SalvarProduto(produto)
    if err != nil {
        return produto, err
    }
    return produto, nil
}

func ListarProdutos() []models.Produto {
    return repositories.ListarProdutos()
}
