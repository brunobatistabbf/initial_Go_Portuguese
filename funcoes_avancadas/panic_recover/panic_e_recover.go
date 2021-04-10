package panic_recover

import "fmt"

func recuperar(){
	if r := recover(); r != nil{
		fmt.Println("Recuperada com sucesso!")
	}

}

func estaAprovado(n1, n2 float64)bool{

	defer recuperar()
	media := (n1 +n2 )/2
	if media > 6 {

		fmt.Println("Aprovado")
		return true
	}else if media < 6 {
		fmt.Println("REPROVADO")
	}

	 panic("A MEDIA É 6")

}
func main()  {
	fmt.Println(estaAprovado(6, 6))
	fmt.Println("Pós execução")
}