package main

import (
	"fmt"
	"time"
)

func main() {

	canal := multiplexar(escrever("Ola mundo"), escrever("Ola mundo 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}
func multiplexar(canalentra, canalentra2 <-chan string) <-chan string {
	canalsaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalentra:
				canalsaida <- mensagem
			case mensagem := <-canalentra2:
				canalsaida <- mensagem
			}
		}
	}()
	return canalsaida

}

func escrever(texto string) <-chan string {
	canal1 := make(chan string)

	go func() {
		for {
			canal1 <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal1
}
