package models

import (
	"fmt"
	"math/rand"
)

type Persona struct {
	Nombre string `default:""`
	Edad   int    `default:"0"`
	NSS    string
	Sexo   string  `default:"H"`
	Peso   float64 `default:"0"`
	Altura float64 `default:"0"`
}

// constructor function
func (persona *Persona) DefaultValues(Nombre string, Edad int, Sexo string, Peso float64, Altura float64) {
	if persona.Nombre == "" {
		persona.Nombre = ""
	} else {
		persona.Nombre = Nombre
	}
	if persona.Sexo == "" {
		persona.Sexo = "H"
	} else {
		persona.Sexo = Sexo
	}
	if persona.Peso == 0 {
		persona.Peso = 0
	} else {
		persona.Peso = Peso
	}
	if persona.Altura == 0 {
		persona.Altura = 0
	} else {
		persona.Altura = Altura
	}
	persona.NSS = persona.generaNSS()
}

//Calcular IMC
func (persona *Persona) CalcularIMC() int {
	imc := persona.Peso / (persona.Altura * persona.Altura)

	if persona.Sexo == "H" {
		if imc < 20 {
			return -1
		} else if imc >= 25 && imc <= 20 {
			return 0
		} else {
			return 1
		}
	} else {
		if imc < 19 {
			return -1
		} else if imc >= 24 && imc <= 19 {
			return 0
		} else {
			return 1
		}
	}
}

//Calcula si es mayor de Edad
func (persona *Persona) EsMayorDeEdad() bool {
	if persona.Edad >= 18 {
		return true
	} else {
		return false
	}
}

//Imprime el valor del objeto
func (persona *Persona) ToString() {
	fmt.Println(persona)
}

func (persona *Persona) generaNSS() string {
	const letras = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 8)
    for i := range b {
        b[i] = letras[rand.Intn(len(letras))]
    }
	return string(b)
}