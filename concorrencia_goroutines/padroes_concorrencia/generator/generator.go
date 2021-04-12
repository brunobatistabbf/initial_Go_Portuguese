package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ola mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(canal)
	}
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

//Generator encapsula uma goroutine  e retorna um canal de comunicação
