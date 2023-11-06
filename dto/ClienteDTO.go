package dto

type ClienteResponse struct {
	Id             string `json:"id"`
	Nome           string `json:"nome"`
	Cidade         string `json:"cidade"`
	Cep            string `json:"cep"`
	DataNascimento string `json:"data_nascimento"`
	Status         string `json:"status"`
}
