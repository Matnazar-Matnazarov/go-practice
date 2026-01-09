package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 5: Struct, Funksiyalar va For Loop (Kengaytirilgan)")
	fmt.Println()

	// 1. STRUCT (STRUKTURA)

	fmt.Println("=== 1. Struct (Struktura) ===")
	fmt.Println()

	// Struct asoslari
	fmt.Println("--- Struct asoslari ---")
	demonstrateStructBasic()
	fmt.Println()

	// Struct fieldlarga kirish
	fmt.Println("--- Struct fieldlarga kirish ---")
	demonstrateStructAccess()
	fmt.Println()

	// Struct pointer
	fmt.Println("--- Struct pointer ---")
	demonstrateStructPointer()
	fmt.Println()

	// Struct embedding
	fmt.Println("--- Struct embedding (inheritance) ---")
	demonstrateStructEmbedding()
	fmt.Println()

	// Nested struct
	fmt.Println("--- Nested struct ---")
	demonstrateNestedStruct()
	fmt.Println()

	// Struct array
	fmt.Println("--- Struct array ---")
	demonstrateStructArray()
	fmt.Println()

	// Struct slice
	fmt.Println("--- Struct slice ---")
	demonstrateStructSlice()
	fmt.Println()

	// Struct map
	fmt.Println("--- Struct map ---")
	demonstrateStructMap()
	fmt.Println()

	// Struct metodlar
	fmt.Println("--- Struct metodlar ---")
	demonstrateStructMethods()
	fmt.Println()

	// Pointer receiver
	fmt.Println("--- Pointer receiver ---")
	demonstratePointerReceiver()
	fmt.Println()

	// Struct solishtirish
	fmt.Println("--- Struct solishtirish ---")
	demonstrateStructComparison()
	fmt.Println()

	// Struct parametr sifatida
	fmt.Println("--- Struct parametr sifatida ---")
	demonstrateStructAsParameter()
	fmt.Println()

	// Struct qaytarish
	fmt.Println("--- Struct qaytarish ---")
	demonstrateStructReturn()
	fmt.Println()

	// Struct filtrlash
	fmt.Println("--- Struct filtrlash ---")
	demonstrateStructFilter()
	fmt.Println()

	// Struct sort
	fmt.Println("--- Struct sort ---")
	demonstrateStructSort()
	fmt.Println()

	// 2. FUNKSIYALAR (KENGAYTIRILGAN)

	fmt.Println("=== 2. Funksiyalar (Kengaytirilgan) ===")
	fmt.Println()

	// Funksiya kompozitsiyasi
	fmt.Println("--- Funksiya kompozitsiyasi ---")
	demonstrateFunctionComposition()
	fmt.Println()

	// Higher-order funksiyalar
	fmt.Println("--- Higher-order funksiyalar ---")
	demonstrateHigherOrderFunction()
	fmt.Println()

	// Closure funksiyalar
	fmt.Println("--- Closure funksiyalar ---")
	demonstrateClosure()
	fmt.Println()

	// Rekursiv funksiyalar
	fmt.Println("--- Rekursiv funksiyalar ---")
	demonstrateRecursiveFunction()
	fmt.Println()

	// Defer funksiyalar
	fmt.Println("--- Defer funksiyalar ---")
	demonstrateDeferFunction()
	fmt.Println()

	// Defer bilan tozalash
	fmt.Println("--- Defer bilan tozalash ---")
	demonstrateDeferCleanup()
	fmt.Println()

	// Defer stack
	fmt.Println("--- Defer stack (LIFO) ---")
	demonstrateDeferStack()
	fmt.Println()

	// Anonim funksiyalar
	fmt.Println("--- Anonim funksiyalar ---")
	demonstrateAnonymousFunction()
	fmt.Println()

	// Variadic funksiyalar (qo'shimcha)
	fmt.Println("--- Variadic funksiyalar (qo'shimcha) ---")
	demonstrateFunctionVariadic()
	fmt.Println()

	// Funksiyani qaytarish
	fmt.Println("--- Funksiyani qaytarish ---")
	demonstrateFunctionAsReturn()
	fmt.Println()

	// Error handling
	fmt.Println("--- Error handling ---")
	demonstrateErrorHandling()
	fmt.Println()

	// Struct bilan funksiyalar
	fmt.Println("--- Struct bilan funksiyalar ---")
	demonstrateFunctionWithStruct()
	fmt.Println()

	// 3. FOR LOOP (KENGAYTIRILGAN)

	fmt.Println("=== 3. For Loop (Kengaytirilgan) ===")
	fmt.Println()

	// Loop optimizatsiyasi
	fmt.Println("--- Loop optimizatsiyasi ---")
	demonstrateLoopOptimization()
	fmt.Println()

	// Shartli loop
	fmt.Println("--- Shartli loop ---")
	demonstrateLoopWithCondition()
	fmt.Println()

	// Erta chiqish
	fmt.Println("--- Erta chiqish (break) ---")
	demonstrateLoopEarlyExit()
	fmt.Println()

	// Elementlarni o'tkazib yuborish
	fmt.Println("--- Elementlarni o'tkazib yuborish (continue) ---")
	demonstrateLoopSkip()
	fmt.Println()

	// Ichki looplar
	fmt.Println("--- Ichki looplar (qo'shimcha) ---")
	demonstrateLoopNested()
	fmt.Println()

	// Slice bilan loop
	fmt.Println("--- Slice bilan loop ---")
	demonstrateLoopWithSlice()
	fmt.Println()

	// Map bilan loop
	fmt.Println("--- Map bilan loop ---")
	demonstrateLoopWithMap()
	fmt.Println()

	// String bilan loop
	fmt.Println("--- String bilan loop ---")
	demonstrateLoopWithString()
	fmt.Println()

	// Pattern chiqarish
	fmt.Println("--- Pattern chiqarish ---")
	demonstrateLoopPattern()
	fmt.Println()

	// Qidirish algoritmlari
	fmt.Println("--- Qidirish algoritmlari ---")
	demonstrateLoopSearch()
	fmt.Println()

	// Sort algoritmlari
	fmt.Println("--- Sort algoritmlari ---")
	demonstrateLoopSort()
	fmt.Println()

	// Struct bilan loop
	fmt.Println("--- Struct bilan loop ---")
	demonstrateLoopWithStruct()
	fmt.Println()

	// Loop performance
	fmt.Println("--- Loop performance ---")
	demonstrateLoopPerformance()
	fmt.Println()

	// Loop bilan error handling
	fmt.Println("--- Loop bilan error handling ---")
	demonstrateLoopWithError()
	fmt.Println()

	// Loop transformatsiya
	fmt.Println("--- Loop transformatsiya (map, filter, reduce) ---")
	demonstrateLoopTransformation()
	fmt.Println()

	// 4. YAKUNIY XULOSA

	fmt.Println("=== Kun 5 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Struct (struktura) - e'lon qilish, fieldlar")
	fmt.Println("  ✓ Struct embedding va nested struct")
	fmt.Println("  ✓ Struct metodlar (value va pointer receiver)")
	fmt.Println("  ✓ Struct array, slice, map")
	fmt.Println("  ✓ Funksiya kompozitsiyasi va higher-order funksiyalar")
	fmt.Println("  ✓ Closure va rekursiv funksiyalar")
	fmt.Println("  ✓ Defer funksiyalar va error handling")
	fmt.Println("  ✓ For loop optimizatsiyasi va kengaytirilgan misollar")
	fmt.Println("  ✓ Qidirish va sort algoritmlari")
	fmt.Println("  ✓ Loop transformatsiya (map, filter, reduce)")
}
