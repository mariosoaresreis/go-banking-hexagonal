package erros

import "net/http"

type ApplicationError struct {
	Codigo   int    `json: ",omitempty"`
	Mensagem string `json: mensagem`
}

func (e ApplicationError) Texto() *ApplicationError {
	return &ApplicationError{
		Mensagem: e.Mensagem,
	}

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
