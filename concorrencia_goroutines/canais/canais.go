package main

import "fmt"

func main() {
	//canal de comunicação para sincronizar as goroutines
	fmt.Println("Arquivo Canal de comunicação")
	canal := make(chan string)
	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim fo programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto

	}
}
