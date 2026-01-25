package main

import "testing"

// BenchmarkAdd - Add funksiyasini benchmark qilish
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// BenchmarkMultiply - Multiply funksiyasini benchmark qilish
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(3, 4)
	}
}

// BenchmarkMax - Max funksiyasini benchmark qilish
func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(10, 20)
	}
}

// BenchmarkFactorial - Factorial funksiyasini benchmark qilish
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

// BenchmarkFibonacci - Fibonacci funksiyasini benchmark qilish
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

// BenchmarkAdd_LargeNumbers - Katta sonlar bilan Add benchmark
func BenchmarkAdd_LargeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1000000, 2000000)
	}
}

// BenchmarkMultiply_LargeNumbers - Katta sonlar bilan Multiply benchmark
func BenchmarkMultiply_LargeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(1000, 2000)
	}
}
