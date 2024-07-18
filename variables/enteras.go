package variables

import "fmt"

func MostrarEnteros() {
	var intComunaso int
	intDe32 := int32(12)
	intDe64 := int64(21)

	intComunaso = 20

	fmt.Printf("Imprimo enteros: \n%d - %d - %d\n", intComunaso, intDe32, intDe64)
	fmt.Printf("Imprimo enteros: \n%x - %x - %x", intComunaso, intDe32, intDe64)

}
