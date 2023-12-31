package main

import (
	"fmt"
)

var employees = map[string]int{
	"Benjamin": 20,
	"Nahuel":   26,
	"Brenda":   19,
	"Darío":    44,
	"Pedro":    30,
}

func mayoresQue(m map[string]int, edad int) int {
	count := 0
	for _, v := range m {
		if v > edad {
			count++
		}
	}
	return count
}

func agregarEmpleado(m map[string]int, nombre string, edad int) {
	m[nombre] = edad
	fmt.Printf("Se agregó a %s a la lista de empleados.\n", nombre)
}

func eliminarEmpleado(m map[string]int, nombre string) {
	delete(m, nombre)
	fmt.Printf("%s fue eliminado de la lista de empleados.\n", nombre)
}

func main() {

	//Imprimir edad de Benjamin
	fmt.Printf("Edad de benjamin %d \n", employees["Benjamin"])

	// Ejemplo de uso de las funciones
	fmt.Println("Cantidad de empleados mayores de 21 años:", mayoresQue(employees, 21))

	agregarEmpleado(employees, "Federico", 25)

	eliminarEmpleado(employees, "Pedro")

	// Mostrar la lista actualizada de empleados
	fmt.Println("\nLista actualizada de empleados:")
	for nombre, edad := range employees {
		fmt.Printf("%s: %d años\n", nombre, edad)
	}
}
