package app

import (
	"GoBanking/domain"
	"GoBanking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	router := mux.NewRouter()
	handlers := ClienteHandlers{service.NewClienteService(domain.NewClienteRepositoryDb(getDbClient()))}
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/clientes", handlers.getAllClientes)
	router.HandleFunc("/clientes/{cliente_id[0-9]+}", handlers.getCliente)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	usuario := os.Getenv("USUARIO")
	pass := os.Getenv("PASS")
	dbUrl := os.Getenv("URL")
	porta := os.Getenv("PORTA")
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
