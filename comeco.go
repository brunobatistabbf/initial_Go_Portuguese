package main
//Sempre use oque está depois da ultima barra (/) ex: axuliar.Escrever
//Pacote externo     //Como baixar um pacote externo
//go get  (e o pacote) Ex: github.com/badoux/checkmail
import (
	"fmt"
	"github.com/badoux/checkmail"
	"source/arrays_e_slices"
	"source/auxiliar"
	"source/estruturas_de_controle"
	"source/funcoes"
	"source/funcoes_avancadas/variaticas"
	"source/heranca"
	"source/operadores"
	"source/ponteiros"
	"source/structs"
	"source/tiposdedados"
	"source/variaveis"
	"source/switch"
	"source/loops"
	"source/funcoes_avancadas/retorno_nomeado"
	)

func main(){
	fmt.Println("Olá mundo")
	auxiliar.Escrever()
	//Checador de email //só um teste
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)

	variaveis.Main()
	tiposdedados.Tipodedados()
	funcoes.Main()
	operadores.Main()
	structs.Main()
	heranca.Main()
	ponteiros.Main()
	arrays_e_slices.Main()
	estruturas_de_controle.Main()
	_switch.Main()
	loops.Main()
	retorno_nomeado.Main()
	variaticas.Main()

}
