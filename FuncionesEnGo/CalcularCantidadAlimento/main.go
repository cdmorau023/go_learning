package main

import "fmt"

const (
	Perro     = "perro"
	Gato      = "gato"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func perro(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 10
}
func gato(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 5
}
func hamster(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 0.250
}
func tarantula(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 0.150
}
func animal(animales string) (func(cantidad int) float64, string) {
	var msg string
	switch animales {
	case Perro:
		return perro, msg
	case Gato:
		return gato, msg
	case Hamster:
		return hamster, msg
	case Tarantula:
		return tarantula, msg
	default:
		msg = "Animal no encontrado"
		return nil, msg
	}

}
func main() {

	animalPerro, msg1 := animal(Perro)
	if msg1 != "" {
		panic(msg1)
	}
	animalGato, msg2 := animal(Gato)
	if msg2 != "" {
		panic(msg1)
	}
	animalHamster, msg3 := animal(Hamster)
	if msg3 != "" {
		panic(msg1)
	}
	animalTarantula, msg4 := animal(Tarantula)
	if msg4 != "" {
		panic(msg1)
	}

	var amount float64
	amount += animalPerro(10)
	amount += animalGato(10)
	amount += animalHamster(10)
	amount += animalTarantula(10)

	fmt.Printf("%.2f kg de comida en total", amount)
	fmt.Printf(msg1, msg2, msg3, msg4)
}
