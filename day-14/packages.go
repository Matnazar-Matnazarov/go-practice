package main

import (
	"fmt"
	"strings"
)


// 1. PACKAGES ASOSLARI


// demonstratePackageDeclaration - Package e'lon qilish
func demonstratePackageDeclaration() {
	fmt.Println("  package main - executable dastur uchun")
	fmt.Println("  package math - library package uchun")
	fmt.Println("  package utils - utility package uchun")
}

// demonstratePackageNaming - Package name qoidalari
func demonstratePackageNaming() {
	fmt.Println("  ✅ To'g'ri: package math, package utils, package mypackage")
	fmt.Println("  ❌ Xato: package Math, package MATH, package my_package (camelCase yaxshiroq)")
	fmt.Println("  Qoida: Package nomi kichik harflar bilan yoziladi")
}


// 2. IMPORT


// demonstrateImport - Import statement
func demonstrateImport() {
	fmt.Println("  import \"fmt\" - fmt package ni import qilish")
	fmt.Println("  import \"math\" - math package ni import qilish")
	fmt.Println("  Qoida: Import - package dan foydalanish uchun")
}

// demonstrateMultipleImports - Multiple imports
func demonstrateMultipleImports() {
	fmt.Println("  Multiple imports:")
	fmt.Println("    import (")
	fmt.Println("        \"fmt\"")
	fmt.Println("        \"math\"")
	fmt.Println("        \"strings\"")
	fmt.Println("    )")
	fmt.Println("  Qoida: Multiple imports - qavs ichida")
}

// demonstrateImportAliases - Import aliases
func demonstrateImportAliases() {
	fmt.Println("  Import aliases:")
	fmt.Println("    import f \"fmt\"  - fmt ni f deb chaqirish")
	fmt.Println("    import m \"math\" - math ni m deb chaqirish")
	fmt.Println("  Qoida: Alias - package nomini o'zgartirish")
}


// 3. EXPORTED VA UNEXPORTED


// ExportedFunction - Exported funksiya (katta harf bilan boshlanadi)
func ExportedFunction() {
	fmt.Println("  ExportedFunction() - boshqa package lardan ishlatish mumkin")
}

// unexportedFunction - Unexported funksiya (kichik harf bilan boshlanadi)
func unexportedFunction() {
	fmt.Println("  unexportedFunction() - faqat shu package ichida ishlatish mumkin")
}

// demonstrateExported - Exported identifiers
func demonstrateExported() {
	fmt.Println("  Exported - katta harf bilan boshlanadi:")
	fmt.Println("    func Add() {}      - ✅ Exported")
	fmt.Println("    var Pi = 3.14     - ✅ Exported")
	fmt.Println("    type Person struct {} - ✅ Exported")
	fmt.Println("  Qoida: Katta harf bilan boshlangan - exported")
	ExportedFunction()
}

// demonstrateUnexported - Unexported identifiers
func demonstrateUnexported() {
	fmt.Println("  Unexported - kichik harf bilan boshlanadi:")
	fmt.Println("    func add() {}     - ❌ Unexported")
	fmt.Println("    var pi = 3.14     - ❌ Unexported")
	fmt.Println("    type person struct {} - ❌ Unexported")
	fmt.Println("  Qoida: Kichik harf bilan boshlangan - unexported")
	unexportedFunction()
}

// Counter - Encapsulation misoli
type Counter struct {
	value int // Unexported field
}

// Increment - Exported method
func (c *Counter) Increment() {
	c.value++
}

// GetValue - Exported method
func (c *Counter) GetValue() int {
	return c.value
}

// increment - Unexported method
func (c *Counter) increment() {
	c.value++
}

// demonstrateEncapsulation - Encapsulation
func demonstrateEncapsulation() {
	counter := &Counter{value: 0}
	fmt.Println("  Encapsulation misoli:")
	fmt.Printf("    Counter value: %d\n", counter.GetValue())
	counter.Increment()
	fmt.Printf("    After Increment(): %d\n", counter.GetValue())
	fmt.Println("  Qoida: Encapsulation - API ni boshqarish")
}


// 4. INIT() FUNKSIYASI


// init() - Package yuklanganda avtomatik chaqiriladi
func init() {
	fmt.Println("  init() funksiyasi chaqirildi (package yuklanganda)")
}

// init2() - Ikkinchi init() funksiyasi
func init2() {
	fmt.Println("  init2() funksiyasi chaqirildi")
}

// demonstrateInitFunction - init() funksiyasi
func demonstrateInitFunction() {
	fmt.Println("  init() funksiyasi:")
	fmt.Println("    - Package yuklanganda avtomatik chaqiriladi")
	fmt.Println("    - main() dan oldin chaqiriladi")
	fmt.Println("    - Bir nechta init() funksiyalari bo'lishi mumkin")
}

// demonstrateMultipleInit - Multiple init() funksiyalari
func demonstrateMultipleInit() {
	fmt.Println("  Multiple init() funksiyalari:")
	fmt.Println("    - Ketma-ket chaqiriladi")
	fmt.Println("    - Har bir faylda init() bo'lishi mumkin")
	fmt.Println("    - Chaqirilish tartibi: import -> init() -> main()")
}


// 5. MODULES


// demonstrateModule - Module e'lon qilish
func demonstrateModule() {
	fmt.Println("  Module yaratish:")
	fmt.Println("    go mod init example.com/myproject")
	fmt.Println("  Qoida: go mod init - yangi module yaratish")
}

// demonstrateGoMod - go.mod fayli
func demonstrateGoMod() {
	fmt.Println("  go.mod fayli:")
	fmt.Println("    module example.com/myproject")
	fmt.Println("    go 1.18")
	fmt.Println("    require (")
	fmt.Println("        github.com/example/package v1.2.3")
	fmt.Println("    )")
	fmt.Println("  Qoida: go.mod - module va qaramliklar")
}


// 6. DEPENDENCIES


// demonstrateDependencies - Dependency management
func demonstrateDependencies() {
	fmt.Println("  Dependency management:")
	fmt.Println("    go get github.com/example/package")
	fmt.Println("    go get github.com/example/package@v1.2.3")
	fmt.Println("    go get github.com/example/package@latest")
	fmt.Println("    go mod tidy - qaramliklarni tozalash")
	fmt.Println("  Qoida: go get - package ni o'rnatish")
}


// 7. MUHIM QOIDALAR


// demonstratePackageNamingRules - Package naming
func demonstratePackageNamingRules() {
	fmt.Println("  Package naming qoidalari:")
	fmt.Println("    ✅ Kichik harflar: package math, package utils")
	fmt.Println("    ✅ Qisqa va tushunarli: package str, package http")
	fmt.Println("    ✅ Underscore yoki camelCase: package my_package yoki package mypackage")
	fmt.Println("    ❌ Katta harflar: package Math")
	fmt.Println("    ❌ Uzun nomlar: package myverylongpackagename")
}

// demonstrateImportOrganization - Import organization
func demonstrateImportOrganization() {
	fmt.Println("  Import organization:")
	fmt.Println("    1. Standard library:")
	fmt.Println("       import \"fmt\"")
	fmt.Println("       import \"math\"")
	fmt.Println("    2. Third-party packages:")
	fmt.Println("       import \"github.com/example/package\"")
	fmt.Println("    3. Local packages:")
	fmt.Println("       import \"example.com/myproject/utils\"")
}

// demonstrateBestPractices - Best practices
func demonstrateBestPractices() {
	fmt.Println("  Best practices:")
	fmt.Println("    ✅ Package nomi - kichik harflar, qisqa")
	fmt.Println("    ✅ Exported - faqat kerakli narsalar")
	fmt.Println("    ✅ init() - kam ishlatish")
	fmt.Println("    ✅ Dependencies - minimal qaramliklar")
	fmt.Println("    ✅ Import organization - tartibga soling")
	fmt.Println("    ✅ Encapsulation - API ni boshqarish")
}


// QO'SHIMCHA MISOL: STRING UTILITIES


// StringUtils - String utility funksiyalari
func StringUtils() {
	text := "hello world"
	fmt.Printf("  Original: %s\n", text)
	fmt.Printf("  Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("  Title: %s\n", strings.Title(text))
	fmt.Printf("  Contains 'world': %t\n", strings.Contains(text, "world"))
}
