package ejercicio_interfaces

import (
	"fmt"
	"github.com/ErMamone/GoDesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	hu.Pensar()

	fmt.Printf("Soy un/a %s, y estoy respirando? TAMBIEN PIENSO?\n", hu.Sexo())
}
