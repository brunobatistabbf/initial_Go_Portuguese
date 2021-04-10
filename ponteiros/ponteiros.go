package ponteiros

import "fmt"

func Main()  {
	fmt.Println("Arquivo de ponteiros")
	var variavel int = 100
	var ponteiro *int   //declaracao de ponteiro

	ponteiro = &variavel  //E comercial para passar o endereço
	fmt.Println(ponteiro) //mostra o endereço
	fmt.Println(*ponteiro) //desreferenciação, para mostrar o valor que está apontado

}
//ponteiro é uma referencia de memoria
