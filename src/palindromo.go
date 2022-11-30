package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string = ""
	textLower := strings.ToLower(text)

	for i := len(textLower) - 1; i >= 0; i-- {
		textReverse += string(textLower[i])
		fmt.Println(string(textLower[i]))
	}

	if textLower == textReverse {
		println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}

func PalindromoMain() {

	// slice := []string{"hola", "que", "hace"}
	// for i, valor := range slice {
	// 	fmt.Println(i, valor)
	// }

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println("hola", i, slice[i])
	// }

	//Palindromo
	isPalindromo("Amor a roma")

}
