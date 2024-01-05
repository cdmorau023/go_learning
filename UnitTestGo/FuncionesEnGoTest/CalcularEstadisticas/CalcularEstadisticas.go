package CalcularEstadisticas

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

//funcion que calcula el valor minimo

func FunMin(values ...uint) (min float64) {
	min = float64(values[0])
	for _, value := range values {
		if float64(value) < min {
			min = float64(value)
		}
	}
	return
}

//Funcion que calcula el valor maximo

func FunMax(values ...uint) (max float64) {
	max = float64(values[0])
	for _, value := range values {
		if float64(value) > max {
			max = float64(value)
		}
	}
	return
}

//Funcion que calcula el promedio

func FunMedia(values ...uint) (media float64) {

	media = 0
	for _, value := range values {
		media += float64(value)
	}
	media /= float64(len(values))
	return
}

//Funcion que recibe un string y devuelve una funcion y un error

func operation(operador string) (func(values ...uint) float64, error) {

	switch operador {
	case minimum:
		return FunMin, nil
	case maximum:
		return FunMax, nil
	case average:
		return FunMedia, nil
	default:
		return nil, errors.New("no se encontró operación")
	}

}

// Funcion principal
func main() {
	var err error
	minFunc, err := operation(minimum)
	if err != nil {
		panic(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("El valor minimo es: %.0f\nEl promedio es: %.3f\nEl valor maximo es: %.0f\n", minValue, averageValue, maxValue)
}
