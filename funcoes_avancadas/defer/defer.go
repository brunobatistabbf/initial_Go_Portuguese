package _defer

import "fmt"

func funcao1()  {
	fmt.Println("Executando a função 1")
}
func funcao2()  {
	fmt.Println("Executando a funcao 2")
}
func main()  {
	defer funcao1()//DEFER = deixar pra depois, adiar
	funcao2()
}
//DEFER pode garantir por exemplo que uma conexão seja fechada