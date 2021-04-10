package tiposdedados

import (
	"errors"
	"fmt"
)

//alias significa um apelido
func Tipodedados()  {

	//int8, int16, int32, int64 existem varios tipos de inteiros, numero conforme os bytes
	//int é de acordo com sua necessidade
	//uint é um int sem sinal //rune significa int32 //uint8  = byte

	var numeroe int =-233
	var numero int16 = 100
	fmt.Println(numero)
	fmt.Println(numeroe)

	//numeros reais //float depende da necessidade //não pode declarar so como float
	var numeroreal1 float32 = 12.34
	var numeroreal2 float64 = 1312.65
	fmt.Println(numeroreal1)
	fmt.Println(numeroreal2)

	//strings
	var str string = "Texto"
	fmt.Println(str)

	//mais do CHAR no Go //Isso declara o numero do caractere na tabela ASCII
	char := 'B'
	fmt.Println(char)

	//BOOLEANO
	var buleana bool = true
	fmt.Println(buleana)

	//Tipo erro // tem que exportar errors para colocar alerta no erro
	var erro error = errors.New("Erro de teste")
	fmt.Println(erro)

}
