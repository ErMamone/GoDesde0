package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDelN() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1")

	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Error cargando el numero, error: " + err.Error())
		} else {
			for i := 1; i < 11; i++ {
				fmt.Printf("Tabla del %d!\n %d x %d = %d\n", num, num, i, num*i)
			}
		}
	}

}
