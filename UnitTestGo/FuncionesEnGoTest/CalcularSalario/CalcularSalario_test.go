package CalcularSalario

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSalarioCategoriaA(t *testing.T) {
	salario := Salario(60, "A")
	expected := 1000.0 // Salario esperado para categoría A y 60 minutos

	assert.Equal(t, expected, salario, "El salario para la categoría A no coincide")
}

func TestSalarioCategoriaB(t *testing.T) {
	salario := Salario(60, "B")
	expected := 1500.0 + 1500.0*0.2
	// Salario esperado para categoría B y 60 minutos

	assert.Equal(t, expected, salario, "El salario para la categoría B no coincide")
}

func TestSalarioCategoriaC(t *testing.T) {
	salario := Salario(60, "C")
	expected := 3000.0 + 3000.0*0.5
	// Salario esperado para categoría C y 60 minutos

	assert.Equal(t, expected, salario, "El salario para la categoría C no coincide")
}
