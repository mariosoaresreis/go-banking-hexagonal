package service

import (
	"github.com/mariosoaresreis/go-banking-hexagonal/domain"
	"github.com/mariosoaresreis/go-banking-hexagonal/dto"
	"github.com/mariosoaresreis/go-banking-hexagonal/erros"
)

type ClienteService interface {
	GetAllClientes(string) ([]dto.ClienteResponse, *erros.ApplicationError)
	GetCliente(id string) (*dto.ClienteResponse, *erros.ApplicationError)
}

type DefaultClienteService struct {
	Repo domain.ClienteRepository
}

func (s DefaultClienteService) GetAllClientes(status string) ([]dto.ClienteResponse, *erros.ApplicationError) {
	clientes, err := s.Repo.FindAll(status)

	if err != nil {
		return nil, err
	}
	response := make([]dto.ClienteResponse, 0)

	for _, c := range clientes {
		response = append(response, c.ToDTO())
	}
	return response, err
}
func (s DefaultClienteService) GetCliente(id string) (*dto.ClienteResponse, *erros.ApplicationError) {
	c, err := s.Repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDTO()
	return &response, nil
}

func NewClienteService(repository domain.ClienteRepository) DefaultClienteService {
	return DefaultClienteService{repository}
}
