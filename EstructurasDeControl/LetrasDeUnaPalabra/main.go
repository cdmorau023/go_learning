package main

import (
	"fmt"
)

func main() {

	var entrada string
	var cantidad int

	for {
		fmt.Scanln(&entrada)

		fmt.Println(entrada)
		cantidad = len(entrada)
		fmt.Println("La cantidad de letras es: ", cantidad)
		for _, letra := range entrada {
			fmt.Printf("%c\n", letra)

		}
	}

}
