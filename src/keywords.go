package main

import "fmt"

func keywords() {

	//Uso de keywords defer, break y continue

	//defer: Se ejecuta esa parte del codigo hasta el ultimo
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {

		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("es un Break")
			break
		}

	}

}
