package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDelN() {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error

	for {
		fmt.Println("Ingrese un numero")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				print("Error cargando el numero " + err.Error() + ", porfavor intente de nuevo: \n")
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 11; i++ {
		fmt.Printf("Tabla del %d!\n %d x %d = %d\n", num, num, i, num*i)
	}
}
