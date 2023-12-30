package main

func main() {

	//Correcto
	var lastName string = "Smith"

	//Incorrecto, debe de escribirse sin comillas ya que es un int
	//var age int = "35"
	var age int = 35

	//Incorrecto, porque la palabra reservada boolean no puede servir como un nombre de variable además false de quererlo declarar como bool debe de escribirse sin comillas
	//boolean := "false"
	dual := false

	//Incorrecto, si la variable declarada es String debe de escribirse cómo String
	var salary string = "45857.90"

	//Correcto
	var firstName string = "Mary"

	println(lastName, age, dual, salary, firstName)
}
