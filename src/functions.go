package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgumento(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleRetun(a int) (c, d int) {
	return a, a * 2
}

//cambiar estas funcion por el main
func functions() {
	normalFunction("Hola Mundo soy wicho")
	tripleArgumento(1, 2, "hola")

	value := returnValue(2)
	println("El valor es: ", value)

	value1, value2 := doubleRetun(2)
	fmt.Println("value 1:", value1, "value2:", value2)
}
