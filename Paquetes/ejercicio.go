package main

type Animal struct {
	nombre    string
	velocidad int
}

func CrearAnimal() Animal {
	return Animal{}
}

func (a *Animal) SetNombre(nombre string) {
	a.nombre = nombre
}

func (a *Animal) SetVelocidad(velocidad int) {
	a.velocidad = velocidad
}

func (a *Animal) GetNombre() string {
	return a.nombre
}

func (a *Animal) GetVelocidad() int {
	return a.velocidad
}
