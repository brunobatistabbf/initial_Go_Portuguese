package main

import (
	"crud_basico/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	//CRUD - CREATE, READ, UPDATE, DELETE
	//mux gerenciamento de funções//conter todas rotas de intereção com banco de dados

	router := mux.NewRouter()
	router.HandleFunc("/usuario", servidor.Criarusuario).Methods(http.MethodPost)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
