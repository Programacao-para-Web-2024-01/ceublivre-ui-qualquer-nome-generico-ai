package models

type Pedido struct {
    ID           string   `json:"id"`
    Cliente      string   `json:"cliente"`
    Produtos     []string `json:"produtos"`
    ValorPedido  float64  `json:"valor_pedido"`
    ValorFrete   float64  `json:"valor_frete"`
    Endereco     string   `json:"endereco"`
    Status       string   `json:"status"`
    DataCriacao  string   `json:"data_criacao"`
}
