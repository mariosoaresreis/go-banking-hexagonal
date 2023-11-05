package domain

import "GoBanking/service"

type ClienteRepositoryStub struct {
	clientes []Cliente
}

func (s ClienteRepositoryStub) FindAll() ([]Cliente, error) {
	return s.clientes, nil
}

func NewClienteRepositoryStub() ClienteRepositoryStub {
	clientes := []Cliente{
		{Id: "101", Nome: "Mario", Cidade: "São Paulo", Cep: "01504-001"},
		{Id: "102", Nome: "Larissa", Cidade: "São Paulo", Cep: "01504-001"},
	}

	return ClienteRepositoryStub{clientes}
}

func NewClienteService(repository ClienteRepository) service.DefaultClienteService {
	return service.DefaultClienteService{Repo: repository}
}
