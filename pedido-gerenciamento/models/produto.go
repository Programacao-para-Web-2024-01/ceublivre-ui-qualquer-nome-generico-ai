package models

type Produto struct {
    ID       string  `json:"id"`
    Nome     string  `json:"nome"`
    Quantidade int   `json:"quantidade"`
    Preco    float64 `json:"preco"`
}
