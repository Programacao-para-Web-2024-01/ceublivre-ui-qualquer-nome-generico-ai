package models

type Devolucao struct {
    ID        string `json:"id"`
    PedidoID  string `json:"pedido_id"`
    Motivo    string `json:"motivo"`
    Status    string `json:"status"`
    DataCriacao string `json:"data_criacao"`
}
