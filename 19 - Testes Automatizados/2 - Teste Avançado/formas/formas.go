package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.Area())
}

type Retangulo struct {
	altura  float64
	largura float64
}
type Circulo struct {
	raio float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {
	r := Retangulo{10, 15}
	EscreverArea(r)

	c := Circulo{10}
	EscreverArea(c)
}
