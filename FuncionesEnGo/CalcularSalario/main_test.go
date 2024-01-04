package main

import (
	"testing"
)

func TestSalarioMenorA50K(t *testing.T) {
	Salario := Salario(25000, "A")
	expected := 1000.0

	if Salario != expected {
		t.Errorf("Para Salario menor a $50,000 se esperaba %.2f pero se obtuvo %.2f", expected, Salario)
	}
}

func TestSalarioMayorA50K(t *testing.T) {
	Salario := Salario(75000, "B")
	expected := 2250.0

	if Salario != expected {
		t.Errorf("Para Salario mayor a $50,000 se esperaba %.2f pero se obtuvo %.2f", expected, Salario)
	}
}

func TestSalarioMayorA150K(t *testing.T) {
	Salario := Salario(180000, "C")
	expected := 9000.0

	if Salario != expected {
		t.Errorf("Para Salario mayor a $150,000 se esperaba %.2f pero se obtuvo %.2f", expected, Salario)
	}
}
