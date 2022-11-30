package mypackage

import "fmt"

//CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

//carprivate is type private, siempre hay que poner comentarios para que no te salga error
// type carPrivate struct {
// 	brand string
// 	year  int
// }

//PrintMessage funcion de accseso publico
func PrintMessage(message string) {
	fmt.Println(message)
}
