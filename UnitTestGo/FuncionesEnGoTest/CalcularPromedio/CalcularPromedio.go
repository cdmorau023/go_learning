package CalcularPromedio

import "fmt"

func promedioNotas(values ...uint) (sumMedia float64, size int) {

	for _, value := range values {
		sumMedia += float64(value)
	}
	size = len(values)
	sumMedia = sumMedia / float64(size)

	return
}

func main() {
	a, b := promedioNotas(4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6)

	fmt.Printf("El promedio es %f de %d estudiantes", a, b)

}
