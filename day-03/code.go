package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 3: Shartli operatorlar, Array va Slice")
	fmt.Println()

	// ============================================
	// 1. SHARTLI OPERATORLAR (CONDITIONAL STATEMENTS)
	// ============================================
	fmt.Println("=== 1. Shartli operatorlar ===")
	fmt.Println()

	// if/else
	fmt.Println("--- if/else ---")
	checkAge(25)
	checkAge(15)
	fmt.Println()

	// if/else if/else zanjiri
	fmt.Println("--- if/else if/else zanjiri ---")
	scores := []int{95, 85, 75, 65, 55}
	for _, score := range scores {
		grade := checkGrade(score)
		fmt.Printf("Ball: %d - Baho: %s\n", score, grade)
	}
	fmt.Println()

	// Murakkab shartlar
	fmt.Println("--- Murakkab shartlar ---")
	checkNumber(10)
	checkNumber(-5)
	checkNumber(0)
	fmt.Println()

	// Juft/toq tekshirish
	fmt.Println("--- Juft/toq tekshirish ---")
	checkEvenOdd(4)
	checkEvenOdd(7)
	fmt.Println()

	// Oraliq tekshirish
	fmt.Println("--- Oraliq tekshirish ---")
	checkRange(5)
	checkRange(15)
	checkRange(25)
	fmt.Println()

	// Bir nechta shartlar
	fmt.Println("--- Bir nechta shartlar ---")
	checkMultipleConditions(20, true)
	checkMultipleConditions(20, false)
	checkMultipleConditions(16, true)
	fmt.Println()

	// Bo'linish tekshirish
	fmt.Println("--- Bo'linish tekshirish ---")
	checkDivisibility(15, 3)
	checkDivisibility(15, 4)
	checkDivisibility(10, 0)
	fmt.Println()

	// Switch operatori
	fmt.Println("--- Switch operatori ---")
	fmt.Println("Hafta kunlari:")
	for i := 1; i <= 8; i++ {
		fmt.Printf("Kun %d: ", i)
		switchBasic(i)
	}
	fmt.Println()

	// Switch - bir nechta case
	fmt.Println("--- Switch - bir nechta case ---")
	fmt.Println("Fasllar:")
	for i := 1; i <= 13; i++ {
		fmt.Printf("Oy %d: ", i)
		switchMultiple(i)
	}
	fmt.Println()

	// Switch expression
	fmt.Println("--- Switch expression ---")
	testScores := []int{95, 85, 75, 65, 45}
	for _, score := range testScores {
		result := switchExpression(score)
		fmt.Printf("Ball: %d - Natija: %s\n", score, result)
	}
	fmt.Println()

	// Switch type
	fmt.Println("--- Switch type ---")
	switchType(42)
	switchType("Go")
	switchType(3.14)
	switchType(true)
	switchType([]int{1, 2, 3})
	fmt.Println()

	// Amaliy misollar
	fmt.Println("--- Amaliy misollar ---")

	// Login tekshirish
	isValid := checkLogin("admin", "12345")
	fmt.Printf("Login tekshirish: %t\n", isValid)

	isValid2 := checkLogin("user", "wrong")
	fmt.Printf("Noto'g'ri login: %t\n", isValid2)
	fmt.Println()

	// Harorat tekshirish
	temperatures := []float64{-5, 10, 20, 30, 40}
	for _, temp := range temperatures {
		result := checkTemperature(temp)
		fmt.Printf("Harorat %.1f°C: %s\n", temp, result)
	}
	fmt.Println()

	// Kabisa yili
	years := []int{2020, 2021, 2024, 2100, 2000}
	for _, year := range years {
		isLeap := checkLeapYear(year)
		fmt.Printf("%d yil: ", year)
		if isLeap {
			fmt.Println("Kabisa yili")
		} else {
			fmt.Println("Oddiy yil")
		}
	}
	fmt.Println()

	// Saylov huquqi
	fmt.Println(checkVotingEligibility(20, true))
	fmt.Println(checkVotingEligibility(16, true))
	fmt.Println(checkVotingEligibility(25, false))
	fmt.Println()

	// ============================================
	// 2. ARRAY (MASSIVLAR)
	// ============================================
	fmt.Println("=== 2. Array (Massivlar) ===")
	fmt.Println()

	// Array asoslari
	fmt.Println("--- Array asoslari ---")
	demonstrateArray()
	fmt.Println()

	// Array literal
	fmt.Println("--- Array literal ---")
	demonstrateArrayLiteral()
	fmt.Println()

	// Array operatsiyalari
	fmt.Println("--- Array operatsiyalari ---")
	demonstrateArrayOperations()
	fmt.Println()

	// Amaliy misol: Array yig'indisi
	fmt.Println("--- Amaliy misol: Array yig'indisi ---")
	arr := [5]int{10, 20, 30, 40, 50}
	sum := sumArray(arr)
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Yig'indi: %d\n", sum)
	fmt.Println()

	// ============================================
	// 3. SLICE (KESMALAR)
	// ============================================
	fmt.Println("=== 3. Slice (Kesmalar) ===")
	fmt.Println()

	// Slice asoslari
	fmt.Println("--- Slice asoslari ---")
	demonstrateSlice()
	fmt.Println()

	// Array dan slice
	fmt.Println("--- Array dan slice yaratish ---")
	demonstrateSliceFromArray()
	fmt.Println()

	// Slice operatsiyalari
	fmt.Println("--- Slice operatsiyalari ---")
	demonstrateSliceOperations()
	fmt.Println()

	// Slice nusxalash
	fmt.Println("--- Slice nusxalash ---")
	demonstrateSliceCopy()
	fmt.Println()

	// Slice dan slice
	fmt.Println("--- Slice dan slice yaratish ---")
	demonstrateSliceSlicing()
	fmt.Println()

	// Slice iteratsiyasi
	fmt.Println("--- Slice iteratsiyasi ---")
	demonstrateSliceIteration()
	fmt.Println()

	// Slice o'zgartirish
	fmt.Println("--- Slice o'zgartirish ---")
	demonstrateSliceModification()
	fmt.Println()

	// Ko'p o'lchamli
	fmt.Println("--- Ko'p o'lchamli array va slice ---")
	demonstrateMultiDimensional()
	fmt.Println()

	// Amaliy misollar
	fmt.Println("--- Amaliy misollar ---")

	// Eng katta son
	numbers := []int{10, 5, 20, 15, 30}
	max := findMax(numbers)
	fmt.Printf("Sonlar: %v\n", numbers)
	fmt.Printf("Eng katta son: %d\n", max)

	// Eng kichik son
	min := findMin(numbers)
	fmt.Printf("Eng kichik son: %d\n", min)

	// Element borligini tekshirish
	contains5 := contains(numbers, 5)
	contains100 := contains(numbers, 100)
	fmt.Printf("5 bor: %t, 100 bor: %t\n", contains5, contains100)

	// Juft sonlarni filtrlash
	evenNumbers := filterEven(numbers)
	fmt.Printf("Juft sonlar: %v\n", evenNumbers)

	// Teskari qilish
	reversed := reverseSlice(numbers)
	fmt.Printf("Teskari: %v\n", reversed)

	// Birlashtirish
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	merged := mergeSlices(slice1, slice2)
	fmt.Printf("Birlashtirilgan: %v\n", merged)
	fmt.Println()

	// ============================================
	// 4. YAKUNIY XULOSA
	// ============================================
	fmt.Println("=== Kun 3 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ if/else shartli operatorlar")
	fmt.Println("  ✓ if/else if/else zanjiri")
	fmt.Println("  ✓ Switch operatori")
	fmt.Println("  ✓ Switch expression va type switch")
	fmt.Println("  ✓ Array (massivlar) - e'lon qilish, ishlatish")
	fmt.Println("  ✓ Slice (kesmalar) - yaratish, append, copy")
	fmt.Println("  ✓ Array va Slice operatsiyalari")
	fmt.Println("  ✓ Ko'p o'lchamli array va slice")
	fmt.Println("  ✓ Amaliy misollar va algoritmlar")
}
