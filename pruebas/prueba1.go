package main

import (
	"fmt"
)

type Commun interface{
	Create()
	//Search()
	Update(marca, modelo, tipo string, anio int)
	Delete(id int)
}

type Car struct{
	Marca	string
	Modelo	string
	Tipo	string
	Anio	int
}

func (c Car) Create() {
	fmt.Printf("Mi carro es un %s %s tipo %s año %v \n", c.Marca, c.Modelo, c.Tipo, c.Anio)
}

func (c Car) Update(marca, modelo, tipo string, anio int) {
	c.Marca = marca
	c.Modelo = modelo
	c.Tipo = tipo
	c.Anio = anio
	fmt.Printf("Se actualizo mi carro, ahora es un %s %s tipo %s año %v \n", c.Marca, c.Modelo, c.Tipo, c.Anio)
}

func (c Car) Delete() {
	fmt.Printf("Se elimino el carro con ID \n")
}

func main(){
	toyota := Car{"Toyota", "Corolla", "Sedan", 2010}
	pontiac := Car{"Pontiac", "Vibe", "Hatchback", 2012}

	toyota.Create()
	toyota.Update("Toyota", "Corolla", "Sedan", 2015)

	pontiac.Create()
	pontiac.Update("Pontiac", "Vibe", "Sedan", 2012)
}
