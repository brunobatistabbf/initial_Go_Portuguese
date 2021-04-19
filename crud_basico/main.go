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

	//CREATE - POST
	//READ - GET
	//UPDATE - PUT
	//DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuario", servidor.Criarusuario).Methods(http.MethodPost)
	router.HandleFunc("/usuario", servidor.Buscarusuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuario{id}", servidor.Buscarusuario).Methods(http.MethodGet)
	router.HandleFunc("usuario/{id}", servidor.Atualizarusuario).Methods(http.MethodPut)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
