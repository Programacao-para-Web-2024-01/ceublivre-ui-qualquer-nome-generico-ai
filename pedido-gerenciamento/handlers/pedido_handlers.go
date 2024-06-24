package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/services"
    "github.com/gorilla/mux"
)

func CriarPedido(w http.ResponseWriter, r *http.Request) {
    var pedido models.Pedido
    if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
        log.Printf("Erro ao decodificar pedido: %v", err)
        http.Error(w, "Erro ao decodificar pedido: "+err.Error(), http.StatusBadRequest)
        return
    }

    log.Printf("Pedido recebido: %+v", pedido)

    pedido, err := services.CriarPedido(pedido)
    if err != nil {
        log.Printf("Erro ao criar pedido: %v", err)
        http.Error(w, "Erro ao criar pedido: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(pedido); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }

    log.Printf("Pedido criado: %+v", pedido)
}

func AtualizarPedido(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var pedido models.Pedido
    if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
        log.Printf("Erro ao decodificar pedido: %v", err)
        http.Error(w, "Erro ao decodificar pedido: "+err.Error(), http.StatusBadRequest)
        return
    }

    pedido, err := services.AtualizarPedido(id, pedido)
    if err != nil {
        log.Printf("Erro ao atualizar pedido: %v", err)
        http.Error(w, "Erro ao atualizar pedido: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(pedido); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}

func ListarPedidos(w http.ResponseWriter, r *http.Request) {
    pedidos := services.ListarPedidos()
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(pedidos); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}

func GerarRelatorio(w http.ResponseWriter, r *http.Request) {
    filtros := r.URL.Query()
    relatorio := services.GerarRelatorio(filtros)
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(relatorio); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}

func AtualizarStatusPedido(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var dadosAtualizacao struct {
        Status string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&dadosAtualizacao); err != nil {
        log.Printf("Erro ao decodificar dados de atualização: %v", err)
        http.Error(w, "Erro ao decodificar dados de atualização: "+err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.AtualizarStatusPedido(id, dadosAtualizacao.Status); err != nil {
        log.Printf("Erro ao atualizar status do pedido: %v", err)
        http.Error(w, "Erro ao atualizar status do pedido: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(map[string]string{"status": "success"}); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}

func AtualizarEnderecoPedido(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var dadosAtualizacao struct {
        Endereco string `json:"endereco"`
    }

    if err := json.NewDecoder(r.Body).Decode(&dadosAtualizacao); err != nil {
        log.Printf("Erro ao decodificar dados de atualização: %v", err)
        http.Error(w, "Erro ao decodificar dados de atualização: "+err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.AtualizarEnderecoPedido(id, dadosAtualizacao.Endereco); err != nil {
        log.Printf("Erro ao atualizar endereço do pedido: %v", err)
        http.Error(w, "Erro ao atualizar endereço do pedido: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(map[string]string{"status": "success"}); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}

func ObterPedido(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    pedido, err := services.ObterPedido(id)
    if err != nil {
        log.Printf("Erro ao obter pedido: %v", err)
        http.Error(w, "Erro ao obter pedido: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(pedido); err != nil {
        log.Printf("Erro ao codificar resposta: %v", err)
        http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
    }
}
