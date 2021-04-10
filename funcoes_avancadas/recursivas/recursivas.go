package recursivas

import "fmt"

func fibonacci(posicao uint)uint  {
	if posicao <= 1{
		return posicao
	}
	return posicao
}

func main()  {
	fmt.Println("Funções recursivas")
	fmt.Println("FIBONACCI")

	posicao := uint(12)
	fmt.Println(fibonacci(posicao))
}