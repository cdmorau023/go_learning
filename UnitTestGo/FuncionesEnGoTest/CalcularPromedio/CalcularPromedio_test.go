package CalcularPromedio

import "testing"

/* El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora nos corresponde realizar los test correspondientes:

Calcular el promedio de las notas de los alumnos.*/

func TestPromedioNotas(t *testing.T) {
	promedio, _ := promedioNotas(4, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6)
	expected := 5.0

	if promedio != expected {
		t.Errorf("El promedio no coincide, se obtuvo %.2f pero se esperaba %.2f", promedio, expected)
	}
}
