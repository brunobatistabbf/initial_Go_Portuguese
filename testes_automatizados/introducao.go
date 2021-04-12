package main

import (
	"fmt"
	"testes_automatizados/enderecos"
)

func main() {
	tipoendereco := enderecos.TipoDeEnderecos("Praça  das rosas")
	fmt.Println(tipoendereco)
}
