package repositories

import (
    "database/sql"
    "pedido-gerenciamento/models"
    "log"
    "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }
}

func SalvarPedido(pedido models.Pedido) error {
    query := `INSERT INTO pedidos (id, data_criacao, status, endereco, cliente, produtos, valor_pedido, valor_frete) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
    _, err := db.Exec(query, pedido.ID, pedido.DataCriacao, pedido.Status, pedido.Endereco, pedido.Cliente, pq.Array(pedido.Produtos), pedido.ValorPedido, pedido.ValorFrete)
    return err
}

func AtualizarPedido(pedido models.Pedido) error {
    query := `UPDATE pedidos SET status = $1, endereco = $2, produtos = $3, valor_pedido = $4, valor_frete = $5 WHERE id = $6`
    _, err := db.Exec(query, pedido.Status, pedido.Endereco, pq.Array(pedido.Produtos), pedido.ValorPedido, pedido.ValorFrete, pedido.ID)
    return err
}

func BuscarPedido(id string) (models.Pedido, error) {
    var pedido models.Pedido
    query := `SELECT id, data_criacao, status, endereco, cliente, produtos, valor_pedido, valor_frete FROM pedidos WHERE id = $1`
    row := db.QueryRow(query, id)
    err := row.Scan(&pedido.ID, &pedido.DataCriacao, &pedido.Status, &pedido.Endereco, &pedido.Cliente, pq.Array(&pedido.Produtos), &pedido.ValorPedido, &pedido.ValorFrete)
    if err != nil {
        return pedido, err
    }
    return pedido, nil
}

func ListarPedidos() []models.Pedido {
    query := `SELECT id, data_criacao, status, endereco, cliente, produtos, valor_pedido, valor_frete FROM pedidos`
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("Erro ao listar pedidos: %v", err)
        return nil
    }
    defer rows.Close()

    var pedidos []models.Pedido
    for rows.Next() {
        var pedido models.Pedido
        err := rows.Scan(&pedido.ID, &pedido.DataCriacao, &pedido.Status, &pedido.Endereco, &pedido.Cliente, pq.Array(&pedido.Produtos), &pedido.ValorPedido, &pedido.ValorFrete)
        if err != nil {
            log.Printf("Erro ao ler pedido: %v", err)
            continue
        }
        pedidos = append(pedidos, pedido)
    }
    return pedidos
}
