package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/services"
)

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
    var produto models.Produto
    err := json.NewDecoder(r.Body).Decode(&produto)
    if err != nil {
        log.Printf("Erro ao decodificar produto: %v", err)
        http.Error(w, "Erro ao decodificar produto: "+err.Error(), http.StatusBadRequest)
        return
    }

    produto, err = services.AdicionarProduto(produto)
    if err != nil {
        log.Printf("Erro ao adicionar produto: %v", err)
        http.Error(w, "Erro ao adicionar produto: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(produto); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
    produtos := services.ListarProdutos()
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(produtos); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}
