package main

import (
	"fmt"
)

func llaveValor() {

	m := make(map[string]int)

	m["Wicho"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Wicho"]
	fmt.Println(value, ok)

}
