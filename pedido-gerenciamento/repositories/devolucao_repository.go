package repositories

import (
    "database/sql"
    "pedido-gerenciamento/models"
    "errors"
    _ "github.com/lib/pq"
    "log"
)

func SalvarDevolucao(devolucao models.Devolucao) error {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return err
    }
    defer db.Close()

    query := `INSERT INTO devolucoes (id, pedido_id, motivo, status, data_criacao) VALUES ($1, $2, $3, $4, $5)`
    _, err = db.Exec(query, devolucao.ID, devolucao.PedidoID, devolucao.Motivo, devolucao.Status, devolucao.DataCriacao)
    if err != nil {
        return err
    }
    return nil
}

func ListarDevolucoes() ([]models.Devolucao, error) {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, pedido_id, motivo, status, data_criacao FROM devolucoes")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var devolucoes []models.Devolucao
    for rows.Next() {
        var devolucao models.Devolucao
        if err := rows.Scan(&devolucao.ID, &devolucao.PedidoID, &devolucao.Motivo, &devolucao.Status, &devolucao.DataCriacao); err != nil {
            log.Fatal(err)
        }
        devolucoes = append(devolucoes, devolucao)
    }
    return devolucoes, nil
}

func BuscarDevolucao(id string) (models.Devolucao, error) {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return models.Devolucao{}, err
    }
    defer db.Close()

    var devolucao models.Devolucao
    query := `SELECT id, pedido_id, motivo, status, data_criacao FROM devolucoes WHERE id=$1`
    err = db.QueryRow(query, id).Scan(&devolucao.ID, &devolucao.PedidoID, &devolucao.Motivo, &devolucao.Status, &devolucao.DataCriacao)
    if err != nil {
        if err == sql.ErrNoRows {
            return devolucao, errors.New("devolução não encontrada")
        }
        return devolucao, err
    }
    return devolucao, nil
}

func AtualizarDevolucao(devolucao models.Devolucao) error {
    db, err := sql.Open("postgres", "user=username password=password dbname=pedido_gerenciamento sslmode=disable")
    if err != nil {
        return err
    }
    defer db.Close()

    query := `UPDATE devolucoes SET status = $1 WHERE id = $2`
    _, err = db.Exec(query, devolucao.Status, devolucao.ID)
    if err != nil {
        return err
    }
    return nil
}
