package operadores

import (
	"fmt"
)

func Main(){
	//Operadores comuns de qualquer linguagem
	//soma +, subtraçao -, divisao /, multiplicação *, restodadivisap %

	//Go é fortemente tipado // Não pode fazer nada com  tipos diferente
	//EX: var numero1  int32 , var numero2 int8 não podem ser somados

	//Operadores de atrubuição
	var variavel1 string = "String"
	varaivel2 := "String"
	fmt.Println(variavel1, varaivel2)


	//Operadores relacionais
	fmt.Println(1 >2)
	fmt.Println(1>=2)
	fmt.Println(1 == 2)
    fmt.Println(1 !=2)

	//Operadores logicos
	fmt.Println(false &&true)
	fmt.Println(true ||false)
	fmt.Println(!true)

	//Operadores Unarios
	numero := 10
	numero++
	numero+= 15    //numero + 15
	numero--
	numero%= 3
	fmt.Println(numero)

	//Operadores Tenario
	var numeroteste = 1
	if numeroteste > 1 {
		fmt.Println("Numero maior que 1")
	}else {
		fmt.Println("Numero igual ou menor que 1")
	}
	fmt.Println(numeroteste)

}
