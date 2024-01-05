package CalcularCantidadAlimento

import "fmt"

// creamos constantes para los animales
const (
	Perro     = "perro"
	Gato      = "gato"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func CantidadAlimento_perro(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 10
}
func CantidadAlimento_gato(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 5
}
func CantidadAlimento_hamster(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 0.250
}
func CantidadAlimento_tarantula(cantidadAnimal int) float64 {
	return float64(cantidadAnimal) * 0.150
}
func CantidadAlimento_animal(animales string) (func(cantidad int) float64, string) {
	var msg string
	switch animales {
	case Perro:
		return CantidadAlimento_perro, msg
	case Gato:
		return CantidadAlimento_gato, msg
	case Hamster:
		return CantidadAlimento_hamster, msg
	case Tarantula:
		return CantidadAlimento_tarantula, msg
	default:
		msg = "Animal no encontrado"
		return nil, msg
	}

}
func main() {

	animalPerro, msg1 := CantidadAlimento_animal(Perro)
	if msg1 != "" {
		panic(msg1)
	}
	animalGato, msg2 := CantidadAlimento_animal(Gato)
	if msg2 != "" {
		panic(msg1)
	}
	animalHamster, msg3 := CantidadAlimento_animal(Hamster)
	if msg3 != "" {
		panic(msg1)
	}
	animalTarantula, msg4 := CantidadAlimento_animal(Tarantula)
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
