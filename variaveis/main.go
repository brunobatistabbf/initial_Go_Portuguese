package variaveis

import (
	"fmt"
	_ "source/auxiliar"
	)

func Main() {
	var variavel1 string = "variavel 1" //As duas são do tipo string
	variavel2 := "variavel 2"           //essas são duas formas de declarar variaveis
	
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 = "variavel 3"              //OUTRA forma de declara variaveis no Go
		variavel4 = "variavel 4"             //Não usa dois pontos desse modo
	)
    fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "variavel5", "variavel6" //Outro metodo de declarar variavel
	fmt.Println(variavel5, variavel6)
}
