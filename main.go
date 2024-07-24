package main

import "github.com/ErMamone/GoDesde0/middleware"

func main() {
	/*
		estado, texto := variables.ConvertirATexto(200)
		fmt.Printf("Estado: %t - Texto: %s", estado, texto)

		os := runtime.GOOS

		if os == "Linux." {
			fmt.Printf("Estas andando en %s perrooooo", os)
		} else {
			fmt.Printf("Estas andando en %s perrooooooo", os)
		}

		if os := runtime.GOOS; os == "Linux." {
			que es esta gronchada go wtf????
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Estas en linux papa")
			break
		case "windows":
			fmt.Println("Estas en Windows papaaa")
		case "darwin":
			fmt.Println("Estas en Darwin")
		default:
			fmt.Printf("Estas en %s\n", os)
		}

		num, res := ejercicios.RespuestaInteger("200")

		fmt.Printf("El numero final es: %d y %s", num, res)

		teclado.IngresarNumeros()

		iteraciones.Iterar()

		archivos.LeerArchivo()

		funciones.Exponenciar(2)

		array_slice.MostrarArrays()

		array_slice.Capacidad()

		maps.MostrarMapas()

		usuarios.AltaUsuario()

		//Basicamente:
		//Se puede tener una interfaz "implementada" a partir de sus metodos en una clase
		//Desde otra clase, se llama una funcion de la interfaz y por parametros recibimos el objeto con la interfaz "implementada"
		//Entonces el metodo puede ser generico desde una funcion para varios objetos de clases que implementen esta inferfaz
		//Ejemplo: metodo compra y tenes las clases credito y debito, entonces saca la plata de diferentes formas, pero sigue siendo desde la misma interfaz
		//En definitiva useless, sirve para reescribir algo que usas 200 veces, no para cosas basicas de 1 vez sola.
		Pedro := new(modelos.Hombre)
		interfaces.HumanosRespirando(Pedro)

		Roberta := new(modelos.Mujer)
		interfaces.HumanosRespirando(Roberta)

		//E GOD, NO SABIA ESTO, FUNCIONO EL REPASO
		panic_defer.DeferMomentus()

		//Parecido al try, catch and finally pero no tan robusto
		panic_defer.EjemploPanic()

		canal := make(chan bool)
		//Go routines (go antes de una llamada a un metodo) = async await de js????
		go rutinas.MiNombreLento("Juan Carlos", canal)
		defer func() { <-canal }()
		fmt.Println("Alooo")

		//<-canal

		//estado := <-canal

		//if estado {
		//	fmt.Println("Estado! es true: ", estado)
		//} else {
		//	fmt.Println("Estado! es false: ", estado)
		//}

		webserver.MiWebServer()
	*/

	middleware.MiMiddleware()

}
