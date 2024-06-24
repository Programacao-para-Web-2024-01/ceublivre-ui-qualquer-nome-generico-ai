package main

import (
    "log"
    "net/http"
    "pedido-gerenciamento/handlers"
    "pedido-gerenciamento/repositories"
    "github.com/gorilla/mux"
)

func main() {
    // Conecte ao banco de dados
    dataSourceName := "user=username password=password dbname=pedido_gerenciamento sslmode=disable"
    repositories.InitDB(dataSourceName)

    router := mux.NewRouter()
    
    // Rotas da API
    router.HandleFunc("/pedidos", handlers.CriarPedido).Methods("POST")
    router.HandleFunc("/pedidos/{id}", handlers.AtualizarPedido).Methods("PUT")
    router.HandleFunc("/pedidos", handlers.ListarPedidos).Methods("GET")
    router.HandleFunc("/relatorio", handlers.GerarRelatorio).Methods("GET")
    router.HandleFunc("/devolucoes", handlers.CriarDevolucao).Methods("POST")
    router.HandleFunc("/devolucoes", handlers.ListarDevolucoes).Methods("GET")
    router.HandleFunc("/devolucoes/{id}", handlers.AtualizarDevolucao).Methods("PUT")
    router.HandleFunc("/produtos", handlers.AdicionarProduto).Methods("POST")
    router.HandleFunc("/produtos", handlers.ListarProdutos).Methods("GET")
    router.HandleFunc("/pedidos/{id}/status", handlers.AtualizarStatusPedido).Methods("PUT")
    router.HandleFunc("/pedidos/{id}/endereco", handlers.AtualizarEnderecoPedido).Methods("PUT")
    router.HandleFunc("/pedidos/{id}", handlers.ObterPedido).Methods("GET")

    // Servindo arquivos estáticos do frontend
    fs := http.FileServer(http.Dir("./frontend"))
    router.PathPrefix("/").Handler(fs)

    log.Println("Servidor rodando na porta 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
