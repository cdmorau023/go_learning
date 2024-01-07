package tickets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	madrugada = "madrugada"
	mañana    = "mañana"
	tarde     = "tarde"
	noche     = "noche"
)

// Función que dada una cadena de texto con una hora cómo "16:00" retorna el entero de la hora
func ObtenerHoraPeriodo(horaTexto string) (int, error) {
	parts := strings.Split(horaTexto, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("formato de hora incorrecto faltan los minutos: %s", horaTexto)
	}

	hora, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("error al convertir la hora a entero: %v", err)
	}

	return hora, nil
}

// Función que dada una hora cómo entero retorna el periodo del día
func ObtenerPeriodo(hora int) (periodo string, err error) {
	switch {
	case hora >= 0 && hora <= 6:
		periodo = madrugada
	case hora >= 6 && hora <= 12:
		periodo = mañana
	case hora >= 12 && hora <= 18:
		periodo = tarde
	case hora >= 18 && hora <= 23:
		periodo = noche
	default:
		err = errors.New("la hora excede el rango de 0 a 23")
	}
	return
}
