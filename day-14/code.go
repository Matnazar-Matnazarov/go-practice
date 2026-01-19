package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 14: Packages va Modules (Paketlar va Modullar)")
	fmt.Println()

	// 1. PACKAGES ASOSLARI

	fmt.Println("=== 1. Packages Asoslari ===")
	fmt.Println()

	// Package e'lon qilish
	fmt.Println("--- Package e'lon qilish ---")
	demonstratePackageDeclaration()
	fmt.Println()

	// Package name qoidalari
	fmt.Println("--- Package name qoidalari ---")
	demonstratePackageNaming()
	fmt.Println()

	// 2. IMPORT

	fmt.Println("=== 2. Import ===")
	fmt.Println()

	// Import statement
	fmt.Println("--- Import statement ---")
	demonstrateImport()
	fmt.Println()

	// Multiple imports
	fmt.Println("--- Multiple imports ---")
	demonstrateMultipleImports()
	fmt.Println()

	// Import aliases
	fmt.Println("--- Import aliases ---")
	demonstrateImportAliases()
	fmt.Println()

	// 3. EXPORTED VA UNEXPORTED

	fmt.Println("=== 3. Exported va Unexported ===")
	fmt.Println()

	// Exported identifiers
	fmt.Println("--- Exported identifiers ---")
	demonstrateExported()
	fmt.Println()

	// Unexported identifiers
	fmt.Println("--- Unexported identifiers ---")
	demonstrateUnexported()
	fmt.Println()

	// Encapsulation
	fmt.Println("--- Encapsulation ---")
	demonstrateEncapsulation()
	fmt.Println()

	// 4. INIT() FUNKSIYASI

	fmt.Println("=== 4. init() Funksiyasi ===")
	fmt.Println()

	// init() funksiyasi
	fmt.Println("--- init() funksiyasi ---")
	demonstrateInitFunction()
	fmt.Println()

	// Multiple init() funksiyalari
	fmt.Println("--- Multiple init() funksiyalari ---")
	demonstrateMultipleInit()
	fmt.Println()

	// 5. MODULES

	fmt.Println("=== 5. Modules ===")
	fmt.Println()

	// Module e'lon qilish
	fmt.Println("--- Module e'lon qilish ---")
	demonstrateModule()
	fmt.Println()

	// go.mod fayli
	fmt.Println("--- go.mod fayli ---")
	demonstrateGoMod()
	fmt.Println()

	// 6. DEPENDENCIES

	fmt.Println("=== 6. Dependencies ===")
	fmt.Println()

	// Dependency management
	fmt.Println("--- Dependency management ---")
	demonstrateDependencies()
	fmt.Println()

	// 7. MUHIM QOIDALAR

	fmt.Println("=== 7. Muhim Qoidalar ===")
	fmt.Println()

	// Package naming
	fmt.Println("--- Package naming ---")
	demonstratePackageNamingRules()
	fmt.Println()

	// Import organization
	fmt.Println("--- Import organization ---")
	demonstrateImportOrganization()
	fmt.Println()

	// Best practices
	fmt.Println("--- Best practices ---")
	demonstrateBestPractices()
	fmt.Println()

	// 8. YAKUNIY XULOSA

	fmt.Println("=== Kun 14 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Package e'lon qilish va package name qoidalari")
	fmt.Println("  ✓ Import statement va multiple imports")
	fmt.Println("  ✓ Import aliases va dot import")
	fmt.Println("  ✓ Exported va unexported identifiers")
	fmt.Println("  ✓ init() funksiyasi va chaqirilish tartibi")
	fmt.Println("  ✓ Modules va go.mod fayli")
	fmt.Println("  ✓ Dependencies va version management")
	fmt.Println("  ✓ Best practices va muhim qoidalar")
}
