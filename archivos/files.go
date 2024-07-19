package archivos

import (
	"bufio"
	"fmt"
	"github.com/ErMamone/GoDesde0/ejercicios"
	"os"
)

var filename string = "./archivos/txt/table.txt"

func GrabarTabla() {
	var text string = ejercicios.TablaDelN()
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	_, err = fmt.Fprintln(file, text)

	if err != nil {
		file.Close()
		return
	}

	file.Close()
}

func SumarTabla() {
	var text string = ejercicios.TablaDelN()

	if Append(filename, text) {
		fmt.Println("Error al concatenar el archivo con el texto")
	}
}

func Append(filename string, text string) bool {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error : " + err.Error())
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error : " + err.Error())
		return false
	}

	file.Close()
	return true
}

func LeerArchivo() {
	archivo, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error al leer archivo")
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		fmt.Println("> " + scanner.Text())
	}
	archivo.Close()
}
