package loops

import (
	"fmt"
	"time"
)

func Main()  {

	i := 0

	for i <  10{
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)


	for j := 0; j < 10; j++{
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	arraynomes := [3]string{"joao","pedro","carlos"}
	for indice, nome := range arraynomes{
		fmt.Println(indice,nome)
	}
	//for  {
	//	fmt.Println("LOOP INFINITO")
	//}

}
//Sempre usa for