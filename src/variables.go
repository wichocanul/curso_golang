package main

import "fmt"

func variables() {

	// Declaracion de constantes
	//=====================================
	// const pi float64 = 3.14
	// const pi2 = 3.14
	// fmt.Println("pi1:", pi)
	// fmt.Print("Pi2:", pi2)

	//Declaracion de variables enteras
	//======================================
	// base := 12
	// var altura int = 14
	// var area int
	// fmt.Println(base, altura, area)

	//Zero values:
	//=====================================
	// var a int
	// var b float64
	// var c string
	// var d bool
	// fmt.Println(a, b, c, d)

	//Calcular area de cuadrado
	//============================================
	// const baseCuadrado = 10
	// areaCuadrado := baseCuadrado * baseCuadrado
	// fmt.Println("Area Cuadrado: ", areaCuadrado)

	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma", result)
	//Resta
	result = y - x
	fmt.Println("Resta", result)

}
