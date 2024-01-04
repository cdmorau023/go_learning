package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElementByIndex_ExistentIndex(t *testing.T) {
	//Arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	expected := 3

	//Act
	result := ElementByIndex(slice, index)

	//Assert
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	// Mostrar información de este test en particular
	t.Logf("TestElementByIndex_ExistentIndex passed")

}

func TestElementByIndex_NonExistentIndex(t *testing.T) {
	//Arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 10
	expected := 0

	//Act
	result := ElementByIndex(slice, index)

	//Assert
	assert.Equal(t, expected, result, "El índice no existe")

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	// Mostrar información de este test en particular
	t.Logf("TestElementByIndex_NonExistentIndex passed")
	//Si se genera un panic en un test, el test falla
}

func TestElementByIndex2_ExistentIndex(t *testing.T) {
	//Arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	expected := 3

	//Act
	result, err := ElementByIndex2(slice, index)

	//Assert
	if err != nil {
		t.Fatalf("Error: %s", err.Error())

	}
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
	// Mostrar información de este test en particular
	t.Logf("TestElementByIndex2_ExistentIndex passed")
}

func TestElementByIndex2_NonExistentIndex(t *testing.T) {
	//Arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 10
	expected := 0

	//Act
	result, err := ElementByIndex2(slice, index)

	//Assert
	assert.NoError(t, err, "El índice no existe")
	if err == nil {
		t.Fatalf("Expected error, got %d", result)
	}
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
	// Mostrar información de este test en particular
	t.Logf("TestElementByIndex2_NonExistentIndex passed")
	//Si se genera un panic en un test, el test falla
}

//diferencia entre assert y require
//assert no para la ejecución del test
//require si para la ejecución del test
