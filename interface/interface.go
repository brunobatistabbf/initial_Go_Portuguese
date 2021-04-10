package _interface

import (
	"fmt"
	"math"
)

type calculoarea interface {
	area()  float64
}

type retangulo struct  {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64  {
	return r.altura * r.largura
}

func (u circulo) area() float64{
	return  math.Pi * (u.raio * u.raio)

}

func main()  {
	fmt.Println("Arquivo de interface")
}
