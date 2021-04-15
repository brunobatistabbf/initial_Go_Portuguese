package HTTP

import (
	"log"
	"net/http"
)

func main() {
	//HTTP é protocolo de comunicação
	//Base da comunicação web
	//Cliente <-> Servidor (Request e response)

	//Rotas
	//URI - identificador do recurso
	//metodo - GET, POST, PUT, DELETE

	http.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola mundo"))

	})

	log.Fatal(http.ListenAndServe(";5000", nil))

}
