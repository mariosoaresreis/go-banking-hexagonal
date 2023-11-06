package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/mariosoaresreis/go-banking-hexagonal/domain"
	"github.com/mariosoaresreis/go-banking-hexagonal/logger"
	"github.com/mariosoaresreis/go-banking-hexagonal/service"
	"log"
	"net/http"
	"os"
	"time"
)

func verificarVariaveisDeAmbiente() {
	envProps := []string{
		"ENDERECO",
		"PORTA",
		"USUARIO",
		"PASS",
		"URL",
		"PORTA_DB",
		"DB",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Variável de ambiente %s não definida. Encerrando aplicação...", k))
		}
	}
}

func Start() {
	verificarVariaveisDeAmbiente()
	router := mux.NewRouter()
	handlers := ClienteHandlers{service.NewClienteService(domain.NewClienteRepositoryDb(getDbClient()))}
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/clientes", handlers.getAllClientes)
	router.HandleFunc("/clientes/{cliente_id[0-9]+}", handlers.getCliente)
	endereco := os.Getenv("ENDERECO")
	porta := os.Getenv("PORTA")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", endereco, porta), router))
}

func getDbClient() *sqlx.DB {
	usuario := os.Getenv("USUARIO")
	pass := os.Getenv("PASS")
	dbUrl := os.Getenv("URL")
	porta := os.Getenv("PORTA_DB")
	db := os.Getenv("DB")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", usuario, pass, dbUrl, porta, db)
	client, err := sqlx.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
