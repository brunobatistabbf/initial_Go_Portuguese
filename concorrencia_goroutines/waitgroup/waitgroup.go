package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	wait.Add(2) //significa 2 go routines q ele tem que esperar

	go func() {
		escrever("Ola mundo")
		wait.Done() //Quando terminar ele tirar do contador quando terminar
	}()

	go func() {
		escrever("Programando em GO")
		wait.Done()
	}()

	wait.Wait() //sincorniza elas

	fmt.Println("Arquivo de concorrencia")
	go escrever("Ola mundo") //goroutine
	escrever("Iae mundo")

}
func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)

	}
}
