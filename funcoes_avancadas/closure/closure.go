package closure

import "fmt"

func closure() func()  {
	texto := "ARMARIO FECHADO"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main()  {
	texto := "Dentro da main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
//é uma função que referencia variaveis fora do corpo