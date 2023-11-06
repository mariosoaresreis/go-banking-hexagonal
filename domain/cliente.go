package domain

import "GoBanking/erros"

type Cliente struct {
	Id             string
	Nome           string
	Cidade         string
	Cep            string
	DataNascimento string
	Status         string
}

type ClienteRepository interface {
	FindAll(status string) ([]Cliente, *erros.ApplicationError)
	FindById(string) (*Cliente, *erros.ApplicationError)
}
