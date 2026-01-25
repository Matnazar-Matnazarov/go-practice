package main

import "errors"

// Add - ikkita sonni qo'shish
func Add(a, b int) int {
	return a + b
}

// Subtract - ikkita sonni ayirish
func Subtract(a, b int) int {
	return a - b
}

// Multiply - ikkita sonni ko'paytirish
func Multiply(a, b int) int {
	return a * b
}

// Divide - ikkita sonni bo'lish
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Max - ikkita sondan kattasini qaytarish
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min - ikkita sondan kichigini qaytarish
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// IsEven - son juft yoki toqligini tekshirish
func IsEven(n int) bool {
	return n%2 == 0
}

// Factorial - faktorial hisoblash
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci - Fibonacci ketma-ketligi
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
