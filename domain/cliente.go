package domain

type Cliente struct {
	Id             string
	Nome           string
	Cidade         string
	Cep            string
	DataNascimento string
	Status         string
}

type ClienteRepository interface {
	FindAll(status string) ([]Cliente, error)
	FindById(string) (*Cliente, error)
}
