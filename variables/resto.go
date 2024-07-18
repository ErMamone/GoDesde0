package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Juancharlo"
	Estado = true
	Sueldo = 1500
	Fecha = time.Now()

	var fechaDesignada string

	fechaDesignada = Fecha.GoString()

	fmt.Printf("Nombre: %s\nEstado: %t\nSueldo: %g\nFecha Designada: %s - %s", Nombre, Estado, Sueldo, Fecha, fechaDesignada)
}

func ConvertirATexto(numero int) (bool, string) {
	text := strconv.Itoa(numero)
	return true, text
}
