package handlers

import (
    "encoding/json"
    "net/http"
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/services"
    "github.com/gorilla/mux"
    "log"
)

func CriarDevolucao(w http.ResponseWriter, r *http.Request) {
    var devolucao models.Devolucao
    if err := json.NewDecoder(r.Body).Decode(&devolucao); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    devolucao, err := services.CriarDevolucao(devolucao)
    if err != nil {
        log.Printf("Erro ao criar devolução: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(devolucao)
}

func ListarDevolucoes(w http.ResponseWriter, r *http.Request) {
    devolucoes, err := services.ListarDevolucoes()
    if err != nil {
        log.Printf("Erro ao listar devoluções: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(devolucoes)
}

func AtualizarDevolucao(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var devolucao models.Devolucao
    if err := json.NewDecoder(r.Body).Decode(&devolucao); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    devolucao, err := services.AtualizarDevolucao(id, devolucao)
    if err != nil {
        log.Printf("Erro ao atualizar devolução: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(devolucao)
}
