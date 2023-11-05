package app

import (
	"GoBanking/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Cliente struct {
	Nome string `json:"nome"`
	CEP  int
}

type ClienteHandlers struct {
	service service.ClienteService
}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")
}
func (ch *ClienteHandlers) getAllClientes(writer http.ResponseWriter, request *http.Request) {
	clientes, _ := ch.service.GetAllClientes()
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(clientes)
}

func (ch *ClienteHandlers) getCliente(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["cliente_id"]
	cliente, erro := ch.service.GetCliente(id)

	if erro != nil {
		if erro == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
		}

		fmt.Fprint(writer, erro.Error())
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(cliente)
	}

}
