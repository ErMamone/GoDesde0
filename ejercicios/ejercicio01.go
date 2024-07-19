package ejercicios

import "strconv"

func RespuestaInteger(dato string) (int, string) {
	var convertir int
	convertir, _ = strconv.Atoi(dato)

	if convertir > 100 {
		return convertir, "Es mayor a 100"
	}

	return convertir, "Es menor a 100"
}
