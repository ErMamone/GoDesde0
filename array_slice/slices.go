package array_slice

import "fmt"

var sliceTablas []int = []int{2, 5, 6}
var arreglo [10]int = [10]int{1, 3, 5, 7, 9, 200, 1, 6, 100}

func MostrarSlices() {

	fmt.Println(sliceTablas)
	fmt.Println(arreglo)

	porcion := arreglo[3:]
	porcion2 := arreglo[:2]
	porcion3 := arreglo[1:4:4]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Println(elementos)
	fmt.Printf("\nLargo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d : slice %d", len(nums), cap(nums), nums)
}
