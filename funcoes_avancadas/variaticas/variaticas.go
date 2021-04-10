package variaticas

import (
	"fmt"
)
//Significa que ela pode receber de nenhum a ( N )  numeros
func soma(numeros ...int)int {
	total :=0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func Main()  {
	fmt.Println("Funcoes Variaticas")
	totaldaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(totaldaSoma)
}
