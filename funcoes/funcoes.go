package funcoes

import "fmt"
							//especifica que tipo de retorno vai dar
func somar(n1 int8, n2 int8)int8{
	return n1 + n2
}
func calculos(n1, n2 int8)(int8, int8)  {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func Main()  {
	soma := somar(10, 20)
	fmt.Println(soma)

	resultadosoma, _ := calculos(10,23) //Use underline (_) quando n quiser chamar a func ou variavel
	fmt.Println(resultadosoma)

	//As funções podem ter mais de um retorno

}