package Erros

import "net/http"

type ApplicationError struct {
	Codigo   int
	Mensagem string
}

func NewNotFoundError(mensagem string) *ApplicationError {
	return &ApplicationError{
		Codigo:   http.StatusNotFound,
		Mensagem: mensagem,
	}
}

func NewUnexpectedError(mensagem string) *ApplicationError {
	return &ApplicationError{
		Codigo:   http.StatusInternalServerError,
		Mensagem: mensagem,
	}
}
