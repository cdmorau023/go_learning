package main

func main() {
	//Incorrecta por empezar con un número
	//var 1firstName string
	var firstName string

	//Correctamente declarada
	var lastName string

	//Incorrecta, primero va el nombre de la variable
	//var int age
	var age int

	//Incorrecta, los nombres de las variables no deben de empezar con números, además de que los ":" solo se usan en caso de instanciar una variable que no había sido declaraba previamente
	//no es una nueva variable y go al ser un lenguaje fuertemente tipado no permite almacenar en un espacio de memoria dedicado a string un int
	//1lastName := 6
	lastName = "6"

	//Incorrecto en go se recomienda usar la nomenclatura camelcase
	//var driver_license = true
	var driverLicense = true

	//Incorrecto no pueden haber nombres con espacios
	//var person height int
	var personHeight int

	//Correcto
	childsNumber := 2

	// En go también es recomendable usar cada variable declarada de alguna forma
	println(firstName, lastName, age, driverLicense, personHeight, childsNumber)
	//6 0 true 0 2

}
