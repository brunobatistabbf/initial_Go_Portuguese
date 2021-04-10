package heranca

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8

}
type estudante struct {
	pessoa               //heranca do Go//Como se tivesse copiado todos os camppos de //pessoa no struct
	curso string
	faculdade string
}

func Main()  {
	fmt.Println("Arquivo herança")
	p1 := pessoa{"Ramelo", "Alves",32}
	fmt.Println(p1)

	e1 := estudante{p1, "medicina ","usp"}
	fmt.Println(e1.nome) //Pode de referenciar ao nome da pessoa
}
