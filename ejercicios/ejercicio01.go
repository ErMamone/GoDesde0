package ejercicios

import "strconv"

func RespuestaInteger(dato string) (int64, string) {
	var convertir int64
	convertir, _ = strconv.ParseInt(dato, 10, 0)

	if convertir > 100 {
		return convertir, "Es mayor a 100"
	}

	return convertir, "Es menor a 100"
}
