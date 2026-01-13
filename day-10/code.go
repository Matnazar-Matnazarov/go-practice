package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 10: Pointers va Interfaces")
	fmt.Println()

	// 1. POINTERS

	fmt.Println("=== 1. Pointers ===")
	fmt.Println()

	// Pointer asoslari
	fmt.Println("--- Pointer asoslari ---")
	demonstratePointerBasic()
	fmt.Println()

	// Pointer orqali o'zgartirish
	fmt.Println("--- Pointer orqali o'zgartirish ---")
	demonstratePointerModify()
	fmt.Println()

	// Nil pointer
	fmt.Println("--- Nil pointer ---")
	demonstratePointerNil()
	fmt.Println()

	// new() funksiyasi
	fmt.Println("--- new() funksiyasi ---")
	demonstratePointerNew()
	fmt.Println()

	// Funksiyaga pointer uzatish
	fmt.Println("--- Funksiyaga pointer uzatish ---")
	demonstratePointerFunction()
	fmt.Println()

	// Pointer qaytarish
	fmt.Println("--- Pointer qaytarish ---")
	demonstratePointerReturn()
	fmt.Println()

	// Struct pointer
	fmt.Println("--- Struct pointer ---")
	demonstratePointerStruct()
	fmt.Println()

	// Array pointer
	fmt.Println("--- Array pointer ---")
	demonstratePointerArray()
	fmt.Println()

	// Slice pointer
	fmt.Println("--- Slice pointer ---")
	demonstratePointerSlice()
	fmt.Println()

	// Pointer taqqoslash
	fmt.Println("--- Pointer taqqoslash ---")
	demonstratePointerComparison()
	fmt.Println()

	// Pointer receiver
	fmt.Println("--- Pointer receiver ---")
	demonstratePointerMethod()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: Swap ---")
	demonstratePointerPractical()
	fmt.Println()

	// 2. INTERFACES

	fmt.Println("=== 2. Interfaces ===")
	fmt.Println()

	// Interface asoslari
	fmt.Println("--- Interface asoslari ---")
	demonstrateInterfaceBasic()
	fmt.Println()

	// Empty interface
	fmt.Println("--- Empty interface ---")
	demonstrateInterfaceEmpty()
	fmt.Println()

	// Type assertion
	fmt.Println("--- Type assertion ---")
	demonstrateInterfaceTypeAssertion()
	fmt.Println()

	// Interface composition
	fmt.Println("--- Interface composition ---")
	demonstrateInterfaceComposition()
	fmt.Println()

	// Polymorphism
	fmt.Println("--- Polymorphism ---")
	demonstrateInterfacePolymorphism()
	fmt.Println()

	// Nil interface
	fmt.Println("--- Nil interface ---")
	demonstrateInterfaceNil()
	fmt.Println()

	// Stringer interface
	fmt.Println("--- Stringer interface ---")
	demonstrateInterfaceStringer()
	fmt.Println()

	// Error interface
	fmt.Println("--- Error interface ---")
	demonstrateInterfaceError()
	fmt.Println()

	// Interface embedding
	fmt.Println("--- Interface embedding ---")
	demonstrateInterfaceEmbedding()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: Sort ---")
	demonstrateInterfacePractical()
	fmt.Println()

	// 3. YAKUNIY XULOSA

	fmt.Println("=== Kun 10 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Pointers - manzil va dereference")
	fmt.Println("  ✓ Pointer operatsiyalari (&, *, new)")
	fmt.Println("  ✓ Pointer orqali o'zgartirish")
	fmt.Println("  ✓ Pointer receiver metodlar")
	fmt.Println("  ✓ Interfaces - polimorfizm mexanizmi")
	fmt.Println("  ✓ Interface implementatsiyasi")
	fmt.Println("  ✓ Type assertion va type switch")
	fmt.Println("  ✓ Interface composition va embedding")
	fmt.Println("  ✓ Stringer va Error interfeyslari")
	fmt.Println("  ✓ Amaliy misollar: Swap, Sort")
}
