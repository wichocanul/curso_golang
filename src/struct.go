package main

import "fmt"

type car struct {
	brand string
	year  int
}

func structFunction() {

	myCar := car{brand: "Ford", year: 2022}
	fmt.Println(myCar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Ferrary"
	fmt.Println(otherCar)

}
