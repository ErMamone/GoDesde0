package panic_defer

import (
	"fmt"
	"log"
)

func DeferMomentus() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("Last message")

	fmt.Println("Segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("Error panicoso \n%v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Valor 1 !!!!!!!!!!!")
	}
}
