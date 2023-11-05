package service

import "GoBanking/domain"

type ClienteService interface {
	GetAllClientes() ([]domain.Cliente, error)
	GetCliente(id string) (*domain.Cliente, error)
}

type DefaultClienteService struct {
	Repo domain.ClienteRepository
}

func (s DefaultClienteService) GetAllClientes() ([]domain.Cliente, error) {
	return s.Repo.FindAll("ativo")
}
func (s DefaultClienteService) GetCliente(id string) (*domain.Cliente, error) {
	return s.Repo.FindById(id)
}

func NewClienteService(repository domain.ClienteRepository) DefaultClienteService {
	return DefaultClienteService{repository}
}
