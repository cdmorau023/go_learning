package internal_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	a := 1
	b := 2
	expected := 3
	result := sum(a, b)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSum2(t *testing.T) {
	a := 1
	b := 2
	expected := 3
	result := sum(a, b)

	assert.Equal(t, expected, result, "La suma no coincide")

}

//https://semaphoreci.com/blog/table-driven-unit-tests-go
