package arrays_e_slices

import (
	"fmt"
	"reflect"
)

func Main()  {
	fmt.Println("ARquivo de arrays e slices")
	var array1 [5]string
	array1[0]= "MHM"
	fmt.Println(array1)


	array2 := [5]string{"alpha"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)


	//SLICE
	slice :=[]int{1,2,2,2,2,2,2,2,2,1,1,2,2,2,2}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))//mostra o tipo

	slice = append(slice, 18)//add elemento ao slice ou array , etc...
	fmt.Println(slice)

 	//Arrays internos
 	slice3 :=make([]float32, 10, 15)
 	fmt.Println(slice3)
 	fmt.Println(len(slice3)) //length
 	fmt.Println(cap(slice3))//capacidade


 	//pode estourar o tamanho limite do slice usando os arrays internos
}


