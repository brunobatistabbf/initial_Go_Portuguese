package funcao_com_ponteiro

import "fmt"

func invertePonteiro(numero int)int  {
	return numero * -1
}

func main()  {
	numero := 20
	numeroINV := invertePonteiro(numero)
	fmt.Println(numeroINV)
}
