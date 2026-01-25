package main

import (
	"errors"
	"testing"
)

// TestAdd - Add funksiyasini test qilish
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

// TestSubtract - Subtract funksiyasini test qilish
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Subtract(5, 3) = %d; want 2", result)
	}
}

// TestMultiply - Multiply funksiyasini test qilish
func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("Multiply(3, 4) = %d; want 12", result)
	}
}

// TestDivide - Divide funksiyasini test qilish
func TestDivide(t *testing.T) {
	// Muvaffaqiyatli holat
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}

	// Xato holat
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error, got nil")
	}
}

// TestMax - Max funksiyasini test qilish
func TestMax(t *testing.T) {
	result := Max(5, 10)
	if result != 10 {
		t.Errorf("Max(5, 10) = %d; want 10", result)
	}

	result = Max(10, 5)
	if result != 10 {
		t.Errorf("Max(10, 5) = %d; want 10", result)
	}
}

// TestMin - Min funksiyasini test qilish
func TestMin(t *testing.T) {
	result := Min(5, 10)
	if result != 5 {
		t.Errorf("Min(5, 10) = %d; want 5", result)
	}

	result = Min(10, 5)
	if result != 5 {
		t.Errorf("Min(10, 5) = %d; want 5", result)
	}
}

// TestIsEven - IsEven funksiyasini test qilish
func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("IsEven(2) = false; want true")
	}

	if IsEven(3) {
		t.Error("IsEven(3) = true; want false")
	}
}

// TestFactorial - Factorial funksiyasini test qilish
func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"factorial of 0", 0, 1},
		{"factorial of 1", 1, 1},
		{"factorial of 5", 5, 120},
		{"factorial of 3", 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Factorial(tt.input)
			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// TestFibonacci - Fibonacci funksiyasini test qilish
func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"fibonacci of 0", 0, 0},
		{"fibonacci of 1", 1, 1},
		{"fibonacci of 2", 2, 1},
		{"fibonacci of 5", 5, 5},
		{"fibonacci of 10", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fibonacci(tt.input)
			if result != tt.expected {
				t.Errorf("Fibonacci(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// TestDivide_ErrorCases - Divide funksiyasining xato holatlarini test qilish
func TestDivide_ErrorCases(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("expected error for division by zero")
	}

	if !errors.Is(err, errors.New("division by zero")) {
		// errors.Is() bilan to'g'ridan-to'g'ri solishtirish ishlamaydi,
		// lekin error message ni tekshirish mumkin
		if err.Error() != "division by zero" {
			t.Errorf("expected 'division by zero', got %v", err)
		}
	}
}
