package CalcularSalario

import "fmt"

func Salario(minutosT float64, categoria string) (salario float64) {
	a := map[string][2]float64{"A": {1000, 0}, "B": {1500, 0.20}, "C": {3000, 0.50}}
	b := a[categoria]
	salario = ((minutosT / 60) * b[0]) * (1 + b[1])

	return
}

func main() {
	fmt.Printf("El salario es %.2f", Salario(60, "C"))
}
