package app

import (
	"GoBanking/service"
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
	status := request.URL.Query().Get("status")
	clientes, erro := ch.service.GetAllClientes(status)

	if erro != nil {
		writeResponse(writer, erro.Codigo, erro.Texto())
	} else {
		writeResponse(writer, http.StatusOK, clientes)
	}
}

func (ch *ClienteHandlers) getCliente(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["cliente_id"]
	cliente, erro := ch.service.GetCliente(id)

	if erro != nil {
		writeResponse(writer, erro.Codigo, erro.Texto())
	} else {
		writeResponse(writer, http.StatusOK, cliente)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		panic(erro)
	}
}
