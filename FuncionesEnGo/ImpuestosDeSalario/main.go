package main

import "fmt"

func impuestos(salario float64) float64 {
	diez := 0.0
	diecisiete := 0.0

	if salario > 50000 {
		diecisiete = salario * 0.17
		if salario > 150000 {
			diez = salario * 0.10
		}
	}
	return diez + diecisiete
}

func main() {

	fmt.Printf(" hola  %.2f adios ", impuestos(1000000))

}
