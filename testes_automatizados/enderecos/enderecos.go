package enderecos

import "strings"

//TipoDeEnderecos verfica o ripo de endereco
func TipoDeEnderecos(endereco string) string {
	tiposvalidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	enderecominusculo := strings.ToLower(endereco)
	primeirapalavra := strings.Split(enderecominusculo, " ")[0]

	enderecotemumtipovalido := false
	for _, tipo := range tiposvalidos {
		if tipo == primeirapalavra {
			enderecotemumtipovalido = true
		}
	}
	if enderecotemumtipovalido {
		return primeirapalavra
	}

	return "Tipo Invalido"
}

//Pesquisar: TDD - Test Driven Development
