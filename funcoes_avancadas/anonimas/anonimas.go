package anonimas

import "fmt"

func main()  {
	func(){
		fmt.Println("Funcão anonima")
	}()
	num := 10
	fmt.Sprintf("Recebido ->%d", num) //sprint para passar dados diferentes (%s, %d ...)

}
