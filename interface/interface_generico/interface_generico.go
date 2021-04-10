package interface_generico

import (
	"fmt"
	"reflect"
)

func generica(interf interface{})  {
	fmt.Println(interf)
}

func main()  {
	generica("String")
	generica(777)
	generica(reflect.Interface)
	generica(4.14)
}
//funciona com tudo