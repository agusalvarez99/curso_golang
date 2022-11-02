package PublicPrivate

type persona struct {
	Nombre   string
	Apellido string
}

type Persona2 struct {
	nombre   string
	apellido string
}

func CrearPersona() persona {
	return persona{}
}

func CrearPersona2() Persona2 {
	return Persona2{}
}

// necesitamos ubicar el puntero del objeto persona2
func (p *Persona2) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona2) SetApellido(apellido string) {
	p.apellido = apellido
}

func (p *Persona2) GetNombre() string {
	return p.nombre
}
func (p *Persona2) GetApellido() string {
	return p.apellido
}
