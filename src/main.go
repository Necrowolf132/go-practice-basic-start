package main

import (
	"fmt"
	"math"
)

func main() {
	//declaration of constants
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi)

	//declaration of variables with self-assigned type, and value

	base := 12

	//declaration of variables with type assigned by the programmer, and value

	var base2 int = 14
	var base3 int

	fmt.Println("Bases:")
	fmt.Println(base, base2, base3)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("Valores Cero:")
	fmt.Println(a, b, c, d)

	//practical example of assignments

	const baseCuadrada int = 10
	areaCuadrada := baseCuadrada * baseCuadrada

	fmt.Println("Area cuadrada:", areaCuadrada)

	//arithmatic operations

	x := 10
	y := 50

	//Suma
	result := x + y

	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//Division
	result = x / y
	fmt.Println("Division:", result)

	//Modulo
	result = x % y
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Auto incremental:", x)

	//Decremental
	x--
	fmt.Println("Auto Decremantal:", x)

	//Calcular el area de un rectangulo, un trapecio, y un circulo

	// Rectangulo

	Base := 5
	Altura := 10
	AreaRectangulo := Base * Altura
	fmt.Println("El Area de una rectangulo con base 5 y altura 10 es:", AreaRectangulo)

	// Trapecio

	Base = 8
	BaseMenor := 4
	Altura = 5
	constante := 2
	AreaTrapecio := (((Base + BaseMenor) * Altura) / constante)
	fmt.Println("El Area de una trapecio con base Mayo 8, base menor 4 y altura 5 es:", AreaTrapecio)

	// Circulo

	radio := 7
	AreaCirculo := (math.Pi * math.Pow(float64(radio), 2))
	fmt.Println("El Area de una ciculo con radio 7 usando la constante de pi es:", AreaCirculo)
}
