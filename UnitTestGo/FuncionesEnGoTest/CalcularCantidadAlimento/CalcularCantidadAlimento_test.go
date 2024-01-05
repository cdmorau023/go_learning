package CalcularCantidadAlimento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCantidadAlimento_gato(t *testing.T) {
	cantidad := CantidadAlimento_gato(4)
	expected := 5.0 * 4.0

	assert.Equal(t, expected, cantidad, "La cantidad de alimento no coincide")
}
func TestCantidadAlimento_perro(t *testing.T) {
	cantidad := CantidadAlimento_perro(4)
	expected := 10.0 * 4.0

	assert.Equal(t, expected, cantidad, "La cantidad de alimento no coincide")

}
func TestCantidadAlimento_hamster(t *testing.T) {
	cantidad := CantidadAlimento_hamster(4)
	expected := 0.250 * 4.0

	assert.Equal(t, expected, cantidad, "La cantidad de alimento no coincide")
}

func TestCantidadAlimento_tarantula(t *testing.T) {
	cantidad := CantidadAlimento_tarantula(4)
	expected := 0.150 * 4.0

	assert.Equal(t, expected, cantidad, "La cantidad de alimento no coincide")
}
