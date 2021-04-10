package maps

import "fmt"

func Main()  {
	fmt.Println("Arquivo maps")
	usuario := map[string]string{ //dentro[tipos da chave] fora e o tipo dos valores
		"nome": "Ana",
		"sobrenome": "Fernandes",

	}
	usuario2 := map[string]map[string]string{
		"nome":{
			"nome":"Gabriela",
			"sobrenome":"Arantes",
		},


	}
		fmt.Println(usuario)
		fmt.Println(usuario2)
	delete(usuario, "nome") //Para deletar

}

