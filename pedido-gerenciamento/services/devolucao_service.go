package services

import (
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/repositories"
    "strconv"
    "time"
    "errors"
    "log"
)

func CriarDevolucao(devolucao models.Devolucao) (models.Devolucao, error) {
    devolucao.ID = "devolucao_" + strconv.FormatInt(time.Now().Unix(), 10)
    devolucao.DataCriacao = time.Now().Format("2006-01-02 15:04:05")
    devolucao.Status = "Pendente"

    err := repositories.SalvarDevolucao(devolucao)
    if err != nil {
        log.Printf("Erro ao salvar devolução: %v", err)
        return devolucao, err
    }

    log.Printf("Devolução criada com sucesso: %+v", devolucao)
    return devolucao, nil
}

func ListarDevolucoes() ([]models.Devolucao, error) {
    devolucoes, err := repositories.ListarDevolucoes()
    if err != nil {
        log.Printf("Erro ao listar devoluções: %v", err)
        return nil, err
    }
    return devolucoes, nil
}

func AtualizarDevolucao(id string, devolucaoAtualizada models.Devolucao) (models.Devolucao, error) {
    devolucao, err := repositories.BuscarDevolucao(id)
    if err != nil {
        log.Printf("Erro ao buscar devolução: %v", err)
        return devolucao, err
    }
    if devolucao.ID == "" {
        return devolucao, errors.New("devolução não encontrada")
    }

    devolucao.Status = devolucaoAtualizada.Status

    err = repositories.AtualizarDevolucao(devolucao)
    if err != nil {
        log.Printf("Erro ao atualizar devolução: %v", err)
        return devolucao, err
    }

    log.Printf("Devolução atualizada com sucesso: %+v", devolucao)
    return devolucao, nil
}
