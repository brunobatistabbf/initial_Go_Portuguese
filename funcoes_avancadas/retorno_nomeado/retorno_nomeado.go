package retorno_nomeado

import"fmt"

func calculoMatematico(n1, n2 int)(soma int, subtracao int)  {
	soma = n1 + n2
	subtracao = n1 - n2
	return //Nesse tipo de função ele ja usa as que foram passadas//Por isso o nome função nomeada
}

func Main()  {
	fmt.Println(
		"Retorno nomeado")
}