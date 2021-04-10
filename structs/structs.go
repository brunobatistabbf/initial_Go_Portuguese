package structs

import (
	"fmt"
)
//sintaxe do struct

type usuario struct {
	nome string
	idade int8
}

func Main()  {
	fmt.Println("Arquivo structs")
	//varios jeitos de usar os structs
	var user usuario
	user.idade = 19
	user.nome = "Arnaldo"
	fmt.Println(user)

	user2 := usuario{"Dagoberto", 45}
	fmt.Println(user2)

	user3 := usuario{nome: "vanderlei"} //no caso a idade = 0
	fmt.Println(user3)








}
//strutcs usa para lidar com vaiors tipos//funciona como uma classe