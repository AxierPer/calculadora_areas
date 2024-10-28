package areas

import "fmt"

type areasFigures interface {
	area() float64
}

type Cuadrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Circulo struct {
	Radio float64
}

func (c Cuadrado) area() float64 {
	var result float64 = c.Base * 2
	return result
}

func (r Rectangulo) area() float64 {
	var result float64 = r.Base * r.Altura
	return result
}

func (t Triangulo) area() float64 {
	var result float64 = (t.Base * t.Altura) / 2
	return result
}

func (c Circulo) area() float64 {
	var result float64 = (3.14159265359 * c.Radio) * 2
	return result
}

func Calcular(a areasFigures) {
	fmt.Println("Area:", a.area(), "cm²")
}
