package main

import "fmt"

func main() {

	var numero int
	_, err := fmt.Scanln(&numero)
	if err != nil || numero < 1 || numero > 12 {
		fmt.Println("Número de mes inválido.")
		return
	}
	//Alternativa con slice/array
	meses := []string{"", "Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	fmt.Printf("%d, %s.\n", numero, meses[numero])

	//Alternativa con maps
	meses2 := map[int]string{
		0:  "Error",
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	fmt.Printf("%d, %s.\n", numero, meses2[numero])

	//Alternativa con switch
	switch numero {
	case 1:
		fmt.Printf("Enero")
		fmt.Printf("%d, %s.\n", numero)

	case 2:
		fmt.Printf("%d, %s.\n", numero, "Febrero")

	case 3:
		fmt.Printf("%d, %s.\n", numero, "Marzo")

	case 4:
		fmt.Printf("%d, %s.\n", numero, "Abril")

	case 5:
		fmt.Printf("%d, %s.\n", numero, "Mayo")

	case 6:
		fmt.Printf("%d, %s.\n", numero, "Junio")

	case 7:
		fmt.Printf("%d, %s.\n", numero, "Julio")

	case 8:
		fmt.Printf("%d, %s.\n", numero, "Agosto")

	case 9:
		fmt.Printf("%d, %s.\n", numero, "Septiembre")

	case 10:
		fmt.Printf("%d, %s.\n", numero, "Octubre")

	case 11:
		fmt.Printf("%d, %s.\n", numero, "Noviembre")

	case 12:
		fmt.Printf("%d, %s.\n", numero, "Diciembre")

	default:
		fmt.Printf("Desconocido")
	}

	//La mejor alternativa en mi criterio serían el map ya que guarda la relación directa entre el número del mes y del mes.
}
