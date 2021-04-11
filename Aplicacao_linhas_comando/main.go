package main

import (
	"Aplicacao_linhas_comando/app"
	"fmt"
	"log"
	"os"
)


func main()  {

	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
