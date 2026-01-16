package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 12: Generics (Umumiy turlar)")
	fmt.Println()

	// 1. GENERICS ASOSLARI

	fmt.Println("=== 1. Generics Asoslari ===")
	fmt.Println()

	// Type parameters
	fmt.Println("--- Type parameters ---")
	demonstrateTypeParameters()
	fmt.Println()

	// Type constraints
	fmt.Println("--- Type constraints ---")
	demonstrateTypeConstraints()
	fmt.Println()

	// Multiple type parameters
	fmt.Println("--- Multiple type parameters ---")
	demonstrateMultipleTypeParameters()
	fmt.Println()

	// 2. TYPE CONSTRAINTS (KENGAYTIRILGAN)

	fmt.Println("=== 2. Type Constraints (Kengaytirilgan) ===")
	fmt.Println()

	// Interface constraints
	fmt.Println("--- Interface constraints ---")
	demonstrateInterfaceConstraints()
	fmt.Println()

	// Method constraints
	fmt.Println("--- Method constraints ---")
	demonstrateMethodConstraints()
	fmt.Println()

	// Custom constraints
	fmt.Println("--- Custom constraints ---")
	demonstrateCustomConstraints()
	fmt.Println()

	// ~ Operator
	fmt.Println("--- ~ Operator ---")
	demonstrateTildeOperator()
	fmt.Println()

	// 3. GENERICS BILAN FUNKSIYALAR

	fmt.Println("=== 3. Generics bilan Funksiyalar ===")
	fmt.Println()

	// Generic funksiyalar
	fmt.Println("--- Generic funksiyalar ---")
	demonstrateGenericFunctions()
	fmt.Println()

	// Generic filter
	fmt.Println("--- Generic filter ---")
	demonstrateGenericFilter()
	fmt.Println()

	// Generic reduce
	fmt.Println("--- Generic reduce ---")
	demonstrateGenericReduce()
	fmt.Println()

	// Generic find
	fmt.Println("--- Generic find ---")
	demonstrateGenericFind()
	fmt.Println()

	// 4. GENERICS BILAN DATA STRUCTURES

	fmt.Println("=== 4. Generics bilan Data Structures ===")
	fmt.Println()

	// Generic stack
	fmt.Println("--- Generic stack ---")
	demonstrateGenericStack()
	fmt.Println()

	// Generic queue
	fmt.Println("--- Generic queue ---")
	demonstrateGenericQueue()
	fmt.Println()

	// Generic map
	fmt.Println("--- Generic map ---")
	demonstrateGenericMap()
	fmt.Println()

	// 5. GENERICS BILAN INTERFACE

	fmt.Println("=== 5. Generics bilan Interface ===")
	fmt.Println()

	// Generic interface
	fmt.Println("--- Generic interface ---")
	demonstrateGenericInterface()
	fmt.Println()

	// Type assertion bilan generics
	fmt.Println("--- Type assertion bilan generics ---")
	demonstrateTypeAssertionWithGenerics()
	fmt.Println()

	// 6. MUHIM QOIDALAR

	fmt.Println("=== 6. Muhim Qoidalar ===")
	fmt.Println()

	// Type inference
	fmt.Println("--- Type inference ---")
	demonstrateTypeInference()
	fmt.Println()

	// Zero value
	fmt.Println("--- Zero value ---")
	demonstrateZeroValue()
	fmt.Println()

	// Pointer types
	fmt.Println("--- Pointer types ---")
	demonstratePointerTypes()
	fmt.Println()

	// 7. YAKUNIY XULOSA

	fmt.Println("=== Kun 12 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Type parameters va type constraints")
	fmt.Println("  ✓ Interface constraints va method constraints")
	fmt.Println("  ✓ Custom constraints va ~ operator")
	fmt.Println("  ✓ Generic funksiyalar (Map, Filter, Reduce, Find)")
	fmt.Println("  ✓ Generic data structures (Stack, Queue, Map)")
	fmt.Println("  ✓ Generic interface va type assertion")
	fmt.Println("  ✓ Type inference va zero value")
}
