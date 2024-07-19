package funciones

import "fmt"

func Exponenciar(numero int) {
	if numero > 1000000 {
		return
	}

	fmt.Println(numero)
	Exponenciar(numero * 2)
}
