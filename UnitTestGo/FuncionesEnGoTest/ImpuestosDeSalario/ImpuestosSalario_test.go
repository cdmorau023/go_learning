package ImpuestosDeSalario

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.

func TestImpuestoSalarioMenorA50K(t *testing.T) {
	Impuesto := Impuestos(25000)

	expected := 0.0

	//imprimir valores con assert
	assert.Equal(t, expected, Impuesto, "El impuesto no coincide, se obtuvo %.2f pero se esperaba %.2f", Impuesto, expected)

}

// Calcular el impuesto en caso de que el empleado gane por encima de $50.000 y menor o igual de 150 000.

func TestImpuestoSalarioMayorA50K(t *testing.T) {
	Impuesto := Impuestos(75000)
	expected := 12750.000000000002

	assert.Equal(t, expected, Impuesto, "El impuesto no coincide")
}

// Calcular el impuesto en caso de que el empleado gane por encima de $150.000.

func TestImpuestoSalarioMayorA150K(t *testing.T) {
	Impuesto := Impuestos(180000)
	expected := 180000.0 * 0.27

	assert.Equal(t, expected, Impuesto, "El impuesto no coincide")
}
