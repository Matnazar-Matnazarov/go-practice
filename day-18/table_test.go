package main

import "testing"

// TestAdd_TableDriven - Table-driven test misoli
func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"zero", 0, 0, 0},
		{"mixed positive negative", -2, 3, 1},
		{"large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply_TableDriven - Multiply uchun table-driven test
func TestMultiply_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"negative numbers", -2, -3, 6},
		{"zero", 5, 0, 0},
		{"mixed", -2, 3, -6},
		{"one", 1, 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMax_TableDriven - Max uchun table-driven test
func TestMax_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"a greater than b", 10, 5, 10},
		{"b greater than a", 5, 10, 10},
		{"equal values", 5, 5, 5},
		{"negative numbers", -5, -10, -5},
		{"mixed", -5, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsEven_TableDriven - IsEven uchun table-driven test
func TestIsEven_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even number", 2, true},
		{"odd number", 3, false},
		{"zero", 0, true},
		{"negative even", -2, true},
		{"negative odd", -3, false},
		{"large even", 1000, true},
		{"large odd", 1001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
