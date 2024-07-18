package main

import (
	"fmt"
	"github.com/ErMamone/GoDesde0/variables"
)

func main() {
	estado, texto := variables.ConvertirATexto(200)

	fmt.Printf("Estado: %t - Texto: %s", estado, texto)
}
