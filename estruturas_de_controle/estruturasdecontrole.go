package estruturas_de_controle
//package main
import "fmt"

func Main()  {
	fmt.Println("Ariquivo Estruturas de controle ")

	var numero = 18
	if numero > 15 {
		fmt.Println("Maior que 15")

	}else {
		fmt.Println("Menor que 15")
	}

	if outronumero := numero; outronumero>0 { //if init//n da pra acessar a variavel fora do if
		fmt.Println("outro numero maior que 0")
	}else if outronumero == 0{
		fmt.Println("Outro numero igua a zero")
	}else{
		fmt.Println("menor que 0")
	}

}