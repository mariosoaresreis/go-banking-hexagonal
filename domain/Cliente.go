package domain

import (
	"github.com/mariosoaresreis/go-banking-hexagonal/dto"
	"github.com/mariosoaresreis/go-banking-hexagonal/erros"
)

type Cliente struct {
	Id             string `db:"id"`
	Nome           string `db:"nome"`
	Cidade         string `db:"cidade"`
	Cep            string `db:"cep"`
	DataNascimento string `db:"data_nascimento""`
	Status         string `db:"status"`
}

func (c Cliente) ToDTO() dto.ClienteResponse {
	return dto.ClienteResponse{
		Id:             c.Id,
		Nome:           c.Nome,
		Cep:            c.Cep,
		Cidade:         c.Cidade,
		DataNascimento: c.DataNascimento,
		Status:         c.Status,
	}
}

type ClienteRepository interface {
	FindAll(status string) ([]Cliente, *erros.ApplicationError)
	FindById(string) (*Cliente, *erros.ApplicationError)
}
