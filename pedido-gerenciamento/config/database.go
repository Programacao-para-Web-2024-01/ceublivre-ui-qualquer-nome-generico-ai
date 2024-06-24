package config

import (
    "database/sql"
    "log"

    _ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

func Connect() {
    var err error
    connStr := "user=username dbname=pedido_gerenciamento sslmode=disable password=password"
    DB, err = sql.Open("pgx", connStr)
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Falha ao verificar conexão ao banco de dados:", err)
    }

    log.Println("Conectado ao banco de dados com sucesso.")
}
