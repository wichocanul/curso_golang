package main

import "fmt"

func arraySlices() {

	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	//len => es el numero de elementos que contiene el array
	//cap => es la maxima capacidad para almacenar elemntos en el array
	fmt.Println(len(array), cap(array))
	//----------------------------------------------------------------------------------------

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	//----------------------------------------------------------------------------------------

	//Metodos en el slice
	//Imprime el indice 0
	fmt.Println(slice[0])
	//Imprime hasta el indice 3
	fmt.Println(slice[:3])
	//Imprime del indice 2 al 4
	fmt.Println(slice[2:4])
	//Imprime del indice 4 en adelante
	fmt.Println(slice[4:])
	//----------------------------------------------------------------------------------------

	//Append
	//NOTA: Solo se pueden agregar elementos a un slice mientras que a un array no se puede ya que es estatico
	//agregar elemento
	slice = append(slice, 7)

	//Append agregar otro slice al slice anterior
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
