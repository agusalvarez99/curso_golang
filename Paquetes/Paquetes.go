package main

//con el import llamamos a los paquetes/librerias
import "fmt"

func main() {
	/*fmt.Println("principl!")
	Paquete1.Prueba()
	Paquete2.Prueba()
	persona := PublicPrivate.CrearPersona()
	persona.Nombre = "Agustin"
	persona.Apellido = "Alvarez"
	fmt.Println(persona)
	persona2 := PublicPrivate.CrearPersona2()
	persona2.SetNombre("Agustin")
	persona2.SetApellido("Alvarez")
	fmt.Println(persona2.GetNombre())
	fmt.Println(persona2.GetApellido())*/

	animal1 := CrearAnimal()
	animal1.SetNombre("Leon")
	animal1.SetVelocidad(50)

	animal2 := CrearAnimal()
	animal2.SetNombre("Guepardo")
	animal2.SetVelocidad(120)

	if animal1.velocidad > animal2.velocidad {
		fmt.Println("El ganador es:", animal1.nombre)
	} else {
		fmt.Println("El ganador es:", animal2.nombre)
	}

}
