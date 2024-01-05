package CalcularEstadisticas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Realizar test para calcular el mínimo de calificaciones.
Realizar test para calcular el máximo de calificaciones.
Realizar test para calcular el promedio de calificaciones.
*/

func TestFunMin(t *testing.T) {
	minimo := FunMin(4, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6)
	expected := 4.0

	assert.Equal(t, expected, minimo, "El mínimo no coincide")
}

func TestFunMax(t *testing.T) {
	maximo := FunMax(4, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6)
	expected := 6.0

	assert.Equal(t, expected, maximo, "El máximo no coincide")
}
func TestFunMedia(t *testing.T) {
	media := FunMedia(4, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6)
	expected := 5.0

	assert.Equal(t, expected, media, "El promedio no coincide")
}
