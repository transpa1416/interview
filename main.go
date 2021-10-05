package main

import (
	"bufio"
	"fmt"
	people "interview/api/models"
	"os"
	"strconv"
	"strings"
)

func main() {	
	//Obtenemos información
	people := getInformation()
	//Se pasa por validación de valores por defecto
	people.DefaultValues(people.Nombre, people.Edad, people.Sexo, people.Peso, people.Altura)
	println("Hola ", people.Nombre)

	switch people.CalcularIMC() {
	case -1: 
		println("Importante: estas por debajo de tu peso ideal")
	case 0: 
		println("Excelente: estas en tu peso ideal")
	case 1: 
		println("Importante: estas por arriba de tu peso ideal")
	}

	if people.EsMayorDeEdad() {
		println("Muy bien eres mayor de edad")
	} else {
		println("Para continuar con la consulta es necesario la presencia de tu padre o tutor, ya que eres menor de edad")
	}
	println("Los datos proporcionados son: ")
	people.ToString()
}

//Función para recuperar los datos del formulario
//return people.Persona Model
func getInformation() people.Persona {
	fmt.Print(`Bienvenido`)
	fmt.Print(`Introduce tu nombre: `)
	reader := bufio.NewReader(os.Stdin)
	nameIn, _ := reader.ReadString('\n')
	fmt.Print(`Introduce tu edad: `)
	ageIn, _ := reader.ReadString('\n')
	fmt.Print(`Introduce tu genero [H:hombre/M:Mujer]: `)
	genderIn, _ := reader.ReadString('\n')
	fmt.Print(`Introduce tu peso: `)
	weightIn, _ := reader.ReadString('\n')
	fmt.Print(`Introduce tu altura: `)
	heightIn, _ := reader.ReadString('\n')
	validation := validateData(ageIn, genderIn, weightIn, heightIn)
	if !validation{
		//println("Error: Datos incorrectos, intente de nuevo por favor")
		getInformation()
	} else {
		name := strings.TrimRight(nameIn, "\r\n")
		ageTxt := strings.TrimRight(ageIn, "\r\n")
		age,_ := strconv.Atoi(ageTxt)
		gender := strings.TrimRight(genderIn, "\r\n")
		weightTxt := strings.TrimRight(weightIn, "\r\n")
		weight,_ := strconv.ParseFloat(weightTxt, 64)
		heightTxt := strings.TrimRight(heightIn, "\r\n")
		height,_ := strconv.ParseFloat(heightTxt, 64)
		return people.Persona{Nombre:name, Edad:age, NSS:"", Sexo:gender, Peso:weight, Altura:height}
	}
	//Cualquier otra cosa inesperada, regresa valores default
	return people.Persona{}
	
}

//Función para validar que los datos del formulario seán validos
//return bandError bool
//en caso de error, imprime los valores erroneos
func validateData(ageIn string, genderIn string, weightIn string, heightIn string) bool {
	var bandError bool = true
	gender := strings.TrimRight(genderIn, "\r\n")
	if gender != "H"{
		if gender != "M" {
			println("El genero ", gender, "no es permitido, solo se permite [H/M]")
			bandError = false
		}
	}
	ageTxt := strings.TrimRight(ageIn, "\r\n")
	_, errAge := strconv.Atoi(ageTxt)
	if errAge != nil {
		println(ageTxt, "La edad debe ser valor numerico")
		bandError = false
	}
	weightTxt := strings.TrimRight(weightIn, "\r\n")
	_, errWeight := strconv.ParseFloat(weightTxt, 64)
	if errWeight != nil {
		println(weightTxt, "El peso debe ser valor numerico")
		bandError = false
	}
	heightTxt := strings.TrimRight(heightIn, "\r\n")
	_, errHeight := strconv.ParseFloat(heightTxt, 64)
	if errHeight != nil {
		println(heightTxt, "La altura debe ser valor numerico")
		bandError = false
	}
	return bandError
}