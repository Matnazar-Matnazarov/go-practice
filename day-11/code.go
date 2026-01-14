package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 11: Funksiyalar va Metodlar (Pointer, Interface)")
	fmt.Println()

	// 1. METODLAR (KENGAYTIRILGAN)

	fmt.Println("=== 1. Metodlar (Kengaytirilgan) ===")
	fmt.Println()

	// Method sets
	fmt.Println("--- Method sets ---")
	demonstrateMethodSets()
	fmt.Println()

	// Pointer receiver qoidalari
	fmt.Println("--- Pointer receiver qoidalari ---")
	demonstratePointerReceiverRules()
	fmt.Println()

	// Value vs Pointer receiver
	fmt.Println("--- Value vs Pointer receiver ---")
	demonstrateValueVsPointerReceiver()
	fmt.Println()

	// Method expressions
	fmt.Println("--- Method expressions ---")
	demonstrateMethodExpressions()
	fmt.Println()

	// Method values
	fmt.Println("--- Method values ---")
	demonstrateMethodValues()
	fmt.Println()

	// Method chaining
	fmt.Println("--- Method chaining ---")
	demonstrateMethodChaining()
	fmt.Println()

	// Interface metodlar
	fmt.Println("--- Interface metodlar ---")
	demonstrateInterfaceMethods()
	fmt.Println()

	// 2. FUNKSIYALAR (POINTER VA INTERFACE)

	fmt.Println("=== 2. Funksiyalar (Pointer va Interface) ===")
	fmt.Println()

	// Funksiya turlari
	fmt.Println("--- Funksiya turlari (Function types) ---")
	demonstrateFunctionTypes()
	fmt.Println()

	// Funksiya qiymatlari
	fmt.Println("--- Funksiya qiymatlari (Function values) ---")
	demonstrateFunctionValues()
	fmt.Println()

	// Closure bilan pointer
	fmt.Println("--- Closure bilan pointer ---")
	demonstrateClosureWithPointer()
	fmt.Println()

	// Higher-order funksiyalar bilan interface
	fmt.Println("--- Higher-order funksiyalar bilan interface ---")
	demonstrateHigherOrderWithInterface()
	fmt.Println()

	// Callback funksiyalar
	fmt.Println("--- Callback funksiyalar ---")
	demonstrateCallbackFunctions()
	fmt.Println()

	// Funksiya parametr sifatida
	fmt.Println("--- Funksiya parametr sifatida ---")
	demonstrateFunctionAsParameter()
	fmt.Println()

	// Funksiya qaytarish
	fmt.Println("--- Funksiya qaytarish ---")
	demonstrateFunctionReturn()
	fmt.Println()

	// 3. INTERFACES (KENGAYTIRILGAN)

	fmt.Println("=== 3. Interfaces (Kengaytirilgan) ===")
	fmt.Println()

	// Interface implementatsiyasi
	fmt.Println("--- Interface implementatsiyasi ---")
	demonstrateInterfaceImplementation()
	fmt.Println()

	// Type assertion bilan metodlar
	fmt.Println("--- Type assertion bilan metodlar ---")
	demonstrateTypeAssertionWithMethods()
	fmt.Println()

	// Interface conversion
	fmt.Println("--- Interface conversion ---")
	demonstrateInterfaceConversion()
	fmt.Println()

	// Interface nil tekshirish
	fmt.Println("--- Interface nil tekshirish ---")
	demonstrateInterfaceNilCheck()
	fmt.Println()

	// Interface composition bilan metodlar
	fmt.Println("--- Interface composition bilan metodlar ---")
	demonstrateInterfaceCompositionWithMethods()
	fmt.Println()

	// 4. AMALIY MISOLLAR

	fmt.Println("=== 4. Amaliy Misollar ===")
	fmt.Println()

	// Builder pattern
	fmt.Println("--- Builder pattern ---")
	demonstrateBuilderPattern()
	fmt.Println()

	// Strategy pattern
	fmt.Println("--- Strategy pattern ---")
	demonstrateStrategyPattern()
	fmt.Println()

	// Observer pattern
	fmt.Println("--- Observer pattern ---")
	demonstrateObserverPattern()
	fmt.Println()

	// 5. YAKUNIY XULOSA

	fmt.Println("=== Kun 11 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Method sets va pointer receiver qoidalari")
	fmt.Println("  ✓ Method expressions va method values")
	fmt.Println("  ✓ Method chaining")
	fmt.Println("  ✓ Funksiya turlari va funksiya qiymatlari")
	fmt.Println("  ✓ Closure bilan pointer ishlash")
	fmt.Println("  ✓ Higher-order funksiyalar bilan interface")
	fmt.Println("  ✓ Interface implementatsiyasi va type assertion")
	fmt.Println("  ✓ Interface composition bilan metodlar")
	fmt.Println("  ✓ Design patterns: Builder, Strategy, Observer")
}
