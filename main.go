package main

import (
	"fmt"
	"project-go/utils/areas"
	"strings"
)

func main() {
	var repuestaUsuario string
	fmt.Println("Calculardor de area de figura geometricas\nDeseas calcular el area de las figuras geometricas de 'circulo', 'cuadrado', 'rectangulo' y 'triangulo'\nSi o No")
	fmt.Scan(&repuestaUsuario)

	if strings.ToLower(repuestaUsuario) == "si" {
		var baseCuadrado float64
		fmt.Println("Digite el size de los lados del cuadrado: ")
		fmt.Scan(&baseCuadrado)
		areas.Calcular(areas.Cuadrado{Base: baseCuadrado})
		if baseCuadrado != 0 {
			var baseRectangulo float64
			var alturaRectangulo float64
			fmt.Println("Digite la base del rectangulo: ")
			fmt.Scan(&baseRectangulo)
			fmt.Println("Digite la altura del rectangulo: ")
			fmt.Scan(&alturaRectangulo)
			areas.Calcular(areas.Rectangulo{Base: baseRectangulo, Altura: alturaRectangulo})
			if alturaRectangulo != 0 {
				var baseTriangulo float64
				var alturaTriangulo float64
				fmt.Println("Digite la base del triangulo: ")
				fmt.Scan(&baseTriangulo)
				fmt.Println("Digite la altura del triangulo: ")
				fmt.Scan(&alturaTriangulo)
				areas.Calcular(areas.Triangulo{Base: baseTriangulo, Altura: alturaTriangulo})
				if alturaTriangulo != 0 {
					var radioCirculo float64
					fmt.Println("Digite el radio del circulo: ")
					fmt.Scan(&radioCirculo)
					areas.Calcular(areas.Circulo{Radio: radioCirculo})
				}
			}
		}
	}
}
