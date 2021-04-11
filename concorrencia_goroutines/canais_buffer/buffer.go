package main

import "fmt"

func main() {
	//buffer expecificar a capacidade do canal
	canal := make(chan string, 2) //<- 2 é a capacidade
	canal <- "Ola mundo"

	//canal sem buffer sempre trava se passar de um valor no canal
	mensagem := <-canal
	fmt.Println(mensagem)

}
