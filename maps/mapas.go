package maps

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "Siempre en fase de grupos"
	paises["Argentina"] = "Siempre campeones"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Rivercito": 100,
		"Boca":      95,
		"Lanus":     50,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("\t%s\t%d\n", equipo, puntaje)
	}

	delete(campeonato, "Lanus")
	fmt.Println(campeonato)

	puntaje, existencia := campeonato["Boca"]
	fmt.Printf("\t puntaje %d \t existe : %t \n", puntaje, existencia)
}
