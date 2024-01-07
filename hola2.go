package main

// Función que recibe varios enteros indefinidos y al final otro entero que es el valor a buscar
func BuscarValor(valor int, values ...int) bool {
	for _, value := range values {
		if value == valor {
			return true
		}
	}
	return false
}

func main() {

}
