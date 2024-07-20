package array_slice

import "fmt"

var tabla [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var matriz [20][30]int

func MostrarArrays() {
	tabla[7] = 33
	tabla[2] = 12

	tabla2 := [10]string{"Robertito", "Juancito"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[5][25] = 15
	fmt.Println(matriz)
}
