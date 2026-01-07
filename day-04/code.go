package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 4: Map (Xarita) va For Loop (Tsikllar)")
	fmt.Println()

	// 1. MAP (XARITA) MA'LUMOT TUZILMASI

	fmt.Println("=== 1. Map (Xarita) ===")
	fmt.Println()

	// Map asoslari
	fmt.Println("--- Map asoslari ---")
	demonstrateMapBasic()
	fmt.Println()

	// Map operatsiyalari
	fmt.Println("--- Map operatsiyalari ---")
	demonstrateMapOperations()
	fmt.Println()

	// Map element borligini tekshirish
	fmt.Println("--- Map element borligini tekshirish ---")
	demonstrateMapCheck()
	fmt.Println()

	// Map iteratsiyasi
	fmt.Println("--- Map iteratsiyasi ---")
	demonstrateMapIteration()
	fmt.Println()

	// Turli tipdagi maplar
	fmt.Println("--- Turli tipdagi maplar ---")
	demonstrateMapTypes()
	fmt.Println()

	// Map nusxalash
	fmt.Println("--- Map nusxalash ---")
	demonstrateMapCopy()
	fmt.Println()

	// Map ni tozalash
	fmt.Println("--- Map ni tozalash ---")
	demonstrateMapClear()
	fmt.Println()

	// Amaliy misollar
	fmt.Println("--- Amaliy misollar ---")

	// So'zlar sonini hisoblash
	text := "Go Go Python Python Java"
	wordCount := countWords(text)
	fmt.Printf("Matn: '%s'\n", text)
	fmt.Println("So'zlar soni:", wordCount)
	fmt.Println()

	// Eng katta qiymat
	scores := map[string]int{
		"Matnazar": 95,
		"Ali":      85,
		"Vali":     90,
	}
	maxKey, maxValue := findMaxValue(scores)
	fmt.Printf("Ballar: %v\n", scores)
	fmt.Printf("Eng yuqori: %s - %d ball\n", maxKey, maxValue)
	fmt.Println()

	// Map larni birlashtirish
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"c": 3, "d": 4}
	merged := mergeMaps(map1, map2)
	fmt.Printf("Map1: %v\n", map1)
	fmt.Printf("Map2: %v\n", map2)
	fmt.Printf("Birlashtirilgan: %v\n", merged)
	fmt.Println()

	// Map ni filtrlash
	ages := map[string]int{
		"Matnazar": 25,
		"Ali":      30,
		"Vali":     18,
		"Bobur":    22,
	}
	filtered := filterMap(ages, 20)
	fmt.Printf("Barcha yoshlar: %v\n", ages)
	fmt.Printf("20 yoshdan katta: %v\n", filtered)
	fmt.Println()

	// Kalitlarni olish
	keys := getKeys(ages)
	fmt.Printf("Kalitlar: %v\n", keys)

	// Qiymatlarni olish
	values := getValues(ages)
	fmt.Printf("Qiymatlar: %v\n", values)
	fmt.Println()

	// Map ni teskari qilish
	original := map[string]int{"bir": 1, "ikki": 2, "uch": 3}
	inverted := invertMap(original)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Teskari: %v\n", inverted)
	fmt.Println()

	// 2. FOR LOOP (TSIKLLAR)

	fmt.Println("=== 2. For Loop (Tsikllar) ===")
	fmt.Println()

	// Oddiy for loop
	fmt.Println("--- Oddiy for loop ---")
	demonstrateForBasic()
	fmt.Println()

	// Shartli for loop
	fmt.Println("--- Shartli for loop ---")
	demonstrateForCondition()
	fmt.Println()

	// Cheksiz for loop
	fmt.Println("--- Cheksiz for loop ---")
	demonstrateForInfinite()
	fmt.Println()

	// For range array
	fmt.Println("--- For range array ---")
	demonstrateForRangeArray()
	fmt.Println()

	// For range slice
	fmt.Println("--- For range slice ---")
	demonstrateForRangeSlice()
	fmt.Println()

	// For range string
	fmt.Println("--- For range string ---")
	demonstrateForRangeString()
	fmt.Println()

	// For range map
	fmt.Println("--- For range map ---")
	demonstrateForRangeMap()
	fmt.Println()

	// Ichki for looplar
	fmt.Println("--- Ichki for looplar ---")
	demonstrateNestedLoops()
	fmt.Println()

	// Break va continue
	fmt.Println("--- Break va continue ---")
	demonstrateBreakContinue()
	fmt.Println()

	// Label bilan break
	fmt.Println("--- Label bilan break ---")
	demonstrateLabeledBreak()
	fmt.Println()

	// Label bilan continue
	fmt.Println("--- Label bilan continue ---")
	demonstrateLabeledContinue()
	fmt.Println()

	// Amaliy misollar
	fmt.Println("--- Amaliy misollar ---")

	// Array yig'indisi
	arr := [5]int{10, 20, 30, 40, 50}
	sum := sumArrayLoop(arr)
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Yig'indi: %d\n", sum)
	fmt.Println()

	// Slice yig'indisi
	numbers := []int{1, 2, 3, 4, 5}
	sum2 := sumSliceLoop(numbers)
	fmt.Printf("Slice: %v\n", numbers)
	fmt.Printf("Yig'indi: %d\n", sum2)
	fmt.Println()

	// Eng katta va kichik son
	max := findMaxLoop(numbers)
	min := findMinLoop(numbers)
	fmt.Printf("Eng katta: %d\n", max)
	fmt.Printf("Eng kichik: %d\n", min)
	fmt.Println()

	// Juft sonlar soni
	evenCount := countEvenLoop(numbers)
	fmt.Printf("Juft sonlar soni: %d\n", evenCount)
	fmt.Println()

	// Array ni teskari qilish
	reversed := reverseArrayLoop(arr)
	fmt.Printf("Original: %v\n", arr)
	fmt.Printf("Teskari: %v\n", reversed)
	fmt.Println()

	// Faktorial
	for i := 1; i <= 5; i++ {
		fact := factorialLoop(i)
		fmt.Printf("%d! = %d\n", i, fact)
	}
	fmt.Println()

	// Fibonacci
	fib := fibonacciLoop(10)
	fmt.Printf("Fibonacci (10 ta): %v\n", fib)
	fmt.Println()

	// Pattern
	fmt.Println("Pattern (5 qator):")
	printPattern(5)
	fmt.Println()

	// Qidirish
	index := searchInSlice(numbers, 3)
	fmt.Printf("Slice: %v\n", numbers)
	fmt.Printf("3 soni index: %d\n", index)
	fmt.Println()

	// Takrorlangan elementlarni olib tashlash
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	unique := removeDuplicates(duplicates)
	fmt.Printf("Takrorlangan: %v\n", duplicates)
	fmt.Printf("Yagona: %v\n", unique)
	fmt.Println()

	// 3. YAKUNIY XULOSA

	fmt.Println("=== Kun 4 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Map (xarita) - e'lon qilish, operatsiyalar")
	fmt.Println("  ✓ Map iteratsiyasi va element tekshirish")
	fmt.Println("  ✓ Turli tipdagi maplar (nested, string to slice)")
	fmt.Println("  ✓ Map nusxalash, tozalash, filtrlash")
	fmt.Println("  ✓ For loop - oddiy, shartli, cheksiz")
	fmt.Println("  ✓ For range - array, slice, string, map")
	fmt.Println("  ✓ Ichki for looplar (nested loops)")
	fmt.Println("  ✓ Break, continue, label")
	fmt.Println("  ✓ Amaliy algoritmlar va misollar")
}
