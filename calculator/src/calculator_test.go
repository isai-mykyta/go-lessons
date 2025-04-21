package calculator

import "testing"

func assertEqual(t *testing.T, expected, result float64) {
	t.Helper()
	t.Errorf("Expected %v but got %v", expected, result)
}

func TestAddition(t *testing.T) {
	result := Addition(3, 2)
	expected := float64(5)

	if result != expected {
		assertEqual(t, expected, result)
	}
}

func TestSubstraction(t *testing.T) {
	result := Subtraction(3, 2)
	expected := float64(1)

	if result != expected {
		assertEqual(t, expected, result)
	}
}

func TestDivision(t *testing.T) {
	result, _ := Division(10, 2)
	expected := float64(5)

	if result != expected {
		assertEqual(t, expected, result)
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := Division(10, 0)

	if err == nil {
		t.Errorf("Expected an error for division by zero, but got nil")
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 5)
	expected := float64(10)

	if result != expected {
		assertEqual(t, expected, result)
	}
}
