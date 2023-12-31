package main

import "fmt"

func prestamo(edad int, empleo bool, antiguedad float32, sueldo float64) {

	if edad > 22 && empleo == true && antiguedad > 1 {
		fmt.Printf("Se le prestará ")
		if sueldo > 100000 {
			fmt.Printf(" sin intereses \n")

		} else {
			fmt.Printf(" con intereses \n")

		}
	} else {
		fmt.Printf("No se le prestará \n")

	}

}
func main() {
	prestamo(45, true, 2, 200000)
	prestamo(15, false, 2, 20000)
	prestamo(45, true, 2, 20000)
	prestamo(45, true, 1, 200000)
	prestamo(45, true, 0.5, 200)
}
