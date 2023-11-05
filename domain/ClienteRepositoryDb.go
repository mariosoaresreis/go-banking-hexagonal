package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type ClienteRepositoryDb struct {
	client *sqlx.DB
}

func (db ClienteRepositoryDb) FindAll(status string) ([]Cliente, error) {
	var err error
	clientes := make([]Cliente, 0)

	if status == "" {
		findAllSql := "select id, nome, cidade, cep, data_nascimento, status from clientes"
		err = db.client.Select(&clientes, findAllSql)
	} else {
		findAllSql := "select id, nome, cidade, cep, data_nascimento, status from clientes where status = ?"
		err = db.client.Select(&clientes, findAllSql, status)
	}

	if err != nil {
		log.Println("Erro ao obter dados da tabela clientes " + err.Error())
		return nil, err
	}

	return clientes, nil
}

func (db ClienteRepositoryDb) FindById(id string) (*Cliente, error) {
	clienteSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Cliente
	err := db.client.Get(&c, clienteSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Cliente não encontrado")
		} else {
			log.Fatal("Error while scanning customer " + err.Error())
			return nil, errors.New("Erro inesperado do banco de dados")
		}
	}
	return &c, nil
}

func NewClienteRepositoryDb(dbClient *sqlx.DB) ClienteRepositoryDb {
	return ClienteRepositoryDb{dbClient}
}
