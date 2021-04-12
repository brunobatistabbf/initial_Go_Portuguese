package enderecos

import "testing"

//Para indentificar o codigo, como um codigo de teste, deve-se usar _test.go
//Teste de unidade // testa um pequena parte do seu codigo
//Funcao do tipo teste //importar testing
//Começa com Test e o nome da funcao que vai testar //recebe um ponteiro do tipo testing
// para executar os testes sem precisar dar o cd use, go test ./...

type cenariodeteste struct {
	enderecoinserido string
	retornoesperado  string
}

//TestTipoDeEndereco is the test
func TestTipoDeEnderecos(t *testing.T) {

	cenariodeteste := []cenariodeteste{
		{"Endereço oloco meu", "ARNALDO CESAR COELHO"},
		{"Endereço oloco meu", "DEANOASDF"},
		{"Endereço oloco meu", "Charlie brown"},
		{"Endereço oloco meu", "Tapioca"},
		{"Endereço oloco meu", "Ola"},
		{"Endereço oloco meu", "Rua antonio nunes"},
	}

	//enderecoprateste := "Rua  Paulista"
	//enderecoesperado := "Rua"
	//enderecorecebido := TipoDeEnderecos(enderecoprateste)

	for _, cenario := range cenariodeteste {
		enderecorecebido := TipoDeEnderecos(cenario.enderecoinserido)
		if enderecorecebido != cenario.retornoesperado {
			t.Error("O tipo inserido está invalido")
		}

	}
	//if enderecorecebido != enderecoesperado{
	//	 t.Error("O tipo recebido não é um endereço valido")
	//}

}

//Para gerar um arquivo de texto contendo o relatorio de cobertura
// go test --coverprofile nomedoarquivo.txt
//para ler esse arquivo
//go tool cover --func=nomedoarquivo.txt
// Mosta oque esta sendo executado em html
//go tool cover --func=nomedoarquivo.txt
