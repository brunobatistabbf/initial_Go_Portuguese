package metodos

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

//metodo salvar usuario
func (u usuario) salvar()  {
	fmt.Println("Salvando usuario")
}

func main()  {
	usuario1 := usuario{"Antonia Fontenele", 86}
	fmt.Println(usuario1)
	usuario1.salvar()
}
