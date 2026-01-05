package main

import (
	"fmt"
	"strings"
)

// greet - oddiy funksiya, parametr va qaytarish qiymati yo'q
func greet() {
	fmt.Println("Salom, Go dasturlash tili!")
}

// greetPerson - bitta parametrli funksiya
func greetPerson(name string) {
	fmt.Printf("Salom, %s!\n", name)
}

// printInfo - ko'p parametrli funksiya
func printInfo(name string, age int, city string) {
	fmt.Printf("Ism: %s, Yosh: %d, Shahar: %s\n", name, age, city)
}

// add - qaytarish qiymatli funksiya, bitta natija
func add(a, b int) int {
	return a + b
}

// multiply - qisqa sintaksis, parametrlar bir xil tip bo'lsa
func multiply(a, b int) int {
	return a * b
}

// divide - float qaytarish qiymati
func divide(a, b float64) float64 {
	return a / b
}

// divideAndRemainder - ko'p qaytarish qiymati, ikkita natija
func divideAndRemainder(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// rectangle - to'rtburchak perimetri va yuzasi
func rectangle(width, height int) (perimeter, area int) {
	perimeter = 2 * (width + height)
	area = width * height
	return
}

// calculate - nomlangan qaytarish qiymatlar, avtomatik return
func calculate(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b
	return
}

// mathOperations - nomlangan qaytarish qiymatlar, ko'paytma va bo'linma
func mathOperations(a, b int) (product, division int) {
	product = a * b
	division = a / b
	return
}

// sum - variadic funksiya, cheksiz parametr
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// concatenate - variadic funksiya, matnlar bilan
func concatenate(texts ...string) string {
	result := ""
	for _, text := range texts {
		result += text
	}
	return result
}

// square - funksiya ichida funksiya chaqirish
func square(x int) int {
	return x * x
}

// sumOfSquares - funksiya ichida funksiya chaqirish
func sumOfSquares(a, b int) int {
	return square(a) + square(b)
}

// complexCalculation - murakkab hisob-kitob
func complexCalculation(x, y int) int {
	sumResult := add(x, y)
	difference := x - y
	return square(sumResult) - square(difference)
}

// CalcFunction - funksiya tipi, funksiyani parametr sifatida uzatish
type CalcFunction func(int, int) int

// calculateWithFunction - funksiya tipi orqali hisoblash
func calculateWithFunction(a, b int, fn CalcFunction) int {
	return fn(a, b)
}

// combineNames - amaliy misol, ism birlashtirish
func combineNames(firstName, lastName string) string {
	return firstName + " " + lastName
}

// calculateAge - amaliy misol, yosh hisoblash
func calculateAge(birthYear int) int {
	currentYear := 2024
	return currentYear - birthYear
}

// circle - amaliy misol, doira yuzasi va perimetri
func circle(radius float64) (area, perimeter float64) {
	const pi = 3.14159
	area = pi * radius * radius
	perimeter = 2 * pi * radius
	return
}

// triangleArea - amaliy misol, uchburchak yuzasi (Geron formulasi)
func triangleArea(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return p * (p - a) * (p - b) * (p - c)
}

// validateEmail - amaliy misol, email tekshirish
func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// formatText - amaliy misol, matnni formatlash
func formatText(text string) string {
	trimmed := strings.TrimSpace(text)
	return strings.Title(trimmed)
}
