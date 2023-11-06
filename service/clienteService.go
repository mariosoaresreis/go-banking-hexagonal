package service

import (
	"GoBanking/domain"
	"GoBanking/erros"
)

type ClienteService interface {
	GetAllClientes(string) ([]domain.Cliente, *erros.ApplicationError)
	GetCliente(id string) (*domain.Cliente, *erros.ApplicationError)
}

type DefaultClienteService struct {
	Repo domain.ClienteRepository
}

func (s DefaultClienteService) GetAllClientes(status string) ([]domain.Cliente, *erros.ApplicationError) {
	return s.Repo.FindAll(status)
}
func (s DefaultClienteService) GetCliente(id string) (*domain.Cliente, *erros.ApplicationError) {
	return s.Repo.FindById(id)
}

func NewClienteService(repository domain.ClienteRepository) DefaultClienteService {
	return DefaultClienteService{repository}
}
