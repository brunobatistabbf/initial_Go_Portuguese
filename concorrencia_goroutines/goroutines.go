package main

import "fmt"

func main() {
	//CONCORRENCIA != PARALELISMO
	//Executa mais de uma tarefa de uma vez, se o processador tiver mais de um nucleo
	fmt.Println("Arquivo de concorrencia")
	go escrever("Ola mundo") //goroutine
	escrever("Iae mundo")

	//Esse GO significa que o codigo não precisa esperar toda execuão da função
	//Pode continuar o codigo sem mesmo ter terminado a execução
}
func escrever(texto string) {
	for {
		fmt.Println(texto)
	}
}
