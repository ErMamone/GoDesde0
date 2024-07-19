package funciones

import "fmt"

func Calculos() {
	//func Calculos() func(a int, b int) bool {
	var num1, num2 int
	num1, num2 = 10, 15

	suma := func(a int, b int) int {
		return a + b
	}

	fmt.Println(suma(num1, num2))

	suma = func(a int, b int) int {
		return a * b * num1
	}

	fmt.Println(suma(num1, num2))
}
