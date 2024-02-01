package Mathops

import (
	"errors"
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(5, 3)
	expected := 8.0

	if result != expected {
		t.Errorf("Addition result was incorrect, got: %f, want: %f", result, expected)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 3)
	expected := 2.0

	if result != expected {
		t.Errorf("Subtraction result was incorrect, got: %f, want: %f", result, expected)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(5, 3)
	expected := 15.0

	if result != expected {
		t.Errorf("Multiplication result was incorrect, got: %f, want: %f", result, expected)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(10, 2)
	expected := 5.0

	if result != expected || err != nil {
		t.Errorf("Division result was incorrect, got: %f, want: %f, error: %v", result, expected, err)
	}
	_, zeroErr := Division(5, 0)
	expectedErr := errors.New("division by zero not allowed")

	if zeroErr == nil || zeroErr.Error() != expectedErr.Error() {
		t.Errorf("Division by zero error was not caught correctly, got: %v, want: %v", zeroErr, expectedErr)
	}
}
func TestPower(t *testing.T) {
	//base sama exponen nya sama2 positif
	result, err := Power(2, 3)
	expected := 8.0

	if result != expected || err != nil {
		t.Errorf("Power result was incorrect, got: %f, want: %f, error: %v", result, expected, err)
	}

	// cek kalau basisnya 0
	result, err = Power(0, 3)
	expected = 0.0

	if result != expected || err != nil {
		t.Errorf("Power result was incorrect, got: %f, want: %f, error: %v", result, expected, err)
	}

	// cek klo basisnya nilainya mines
	result, err = Power(-2, 4)
	expected = 16.0

	if result != expected || err != nil {
		t.Errorf("Power result was incorrect, got: %f, want: %f, error: %v", result, expected, err)
	}

	// cek klo pangkatnya mines
	_, err = Power(0, -2)
	expectedErr := errors.New("undefined result: zero to the power of a negative number")

	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Power by zero to a negative number error was not caught correctly, got: %v, want: %v", err, expectedErr)
	}

	// cek klo basis nya mines dan juga pangkatnya bukan int
	_, err = Power(-2, 0.5)
	expectedErr = errors.New("undefined result: negative base to a non-integer power")

	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Power by negative base to a non-integer number error was not caught correctly, got: %v, want: %v", err, expectedErr)
	}
}