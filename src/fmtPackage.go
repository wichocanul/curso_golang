package main

import "fmt"

func fmtPackage() {

	//declaracion de variables
	helloMessage := "Hello"
	worldessage := "World"

	//Println
	fmt.Println(helloMessage, worldessage)
	//fmt.Println(helloMessage, worldessage)

	//Printf
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	//fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos %T", cursos)

}
