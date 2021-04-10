package _switch

import "fmt"
//sintaxe switch
func diadasemana(numero int) string  {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabadoouuu"
	default:
		return "Numero invalido"
	}
}

//fallthrough é usado nos switch para jogar no proximo a função
func Main()  {
	fmt.Println("Arquivo switch")
}
