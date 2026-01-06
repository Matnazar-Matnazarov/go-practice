package main

import "fmt"

// demonstrateArray - array asoslari
func demonstrateArray() {
	// Array e'lon qilish
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array elementlari:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}
}

// demonstrateArrayLiteral - array literal sintaksisi
func demonstrateArrayLiteral() {
	// Qisqa sintaksis
	colors := [3]string{"Qizil", "Yashil", "Ko'k"}
	fmt.Println("Ranglar:", colors)

	// Avtomatik uzunlik
	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Sonlar:", numbers)
	fmt.Printf("Array uzunligi: %d\n", len(numbers))
}

// demonstrateArrayOperations - array operatsiyalari
func demonstrateArrayOperations() {
	arr := [5]int{10, 20, 30, 40, 50}

	// Elementlarga kirish
	fmt.Printf("Birinchi element: %d\n", arr[0])
	fmt.Printf("Oxirgi element: %d\n", arr[len(arr)-1])

	// Elementni o'zgartirish
	arr[2] = 99
	fmt.Println("O'zgartirilgan array:", arr)

	// Array uzunligi
	fmt.Printf("Array uzunligi: %d\n", len(arr))
}

// demonstrateSlice - slice asoslari
func demonstrateSlice() {
	// Slice e'lon qilish
	var numbers []int
	fmt.Println("Bo'sh slice:", numbers)
	fmt.Printf("Slice uzunligi: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Slice literal
	numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)
}

// demonstrateSliceFromArray - array dan slice yaratish
func demonstrateSliceFromArray() {
	arr := [5]int{10, 20, 30, 40, 50}

	// Array dan slice
	slice1 := arr[1:4] // [20, 30, 40]
	fmt.Println("Slice [1:4]:", slice1)

	// Boshidan
	slice2 := arr[:3] // [10, 20, 30]
	fmt.Println("Slice [:3]:", slice2)

	// Oxirigacha
	slice3 := arr[2:] // [30, 40, 50]
	fmt.Println("Slice [2:]:", slice3)

	// Barcha elementlar
	slice4 := arr[:] // [10, 20, 30, 40, 50]
	fmt.Println("Slice [:]:", slice4)
}

// demonstrateSliceOperations - slice operatsiyalari
func demonstrateSliceOperations() {
	// make() bilan slice yaratish
	numbers := make([]int, 3, 5) // uzunlik: 3, capacity: 5
	fmt.Printf("Slice: %v, Uzunlik: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

	// append() - element qo'shish
	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	fmt.Printf("Append qilingan: %v, Uzunlik: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

	// Bir nechta element qo'shish
	numbers = append(numbers, 30, 40, 50)
	fmt.Printf("Ko'p element qo'shilgan: %v\n", numbers)
}

// demonstrateSliceCopy - slice nusxalash
func demonstrateSliceCopy() {
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// O'zgartirish
	copied[0] = 99
	fmt.Println("O'zgartirilgan copied:", copied)
	fmt.Println("Original o'zgarmadi:", original)
}

// demonstrateSliceSlicing - slice dan slice yaratish
func demonstrateSliceSlicing() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// Slice dan slice
	slice1 := numbers[2:6]
	fmt.Println("Slice [2:6]:", slice1)

	// Capacity va length
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice1), cap(slice1))
}

// demonstrateSliceIteration - slice iteratsiyasi
func demonstrateSliceIteration() {
	numbers := []int{10, 20, 30, 40, 50}

	// For loop
	fmt.Println("For loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	// Range loop
	fmt.Println("\nRange loop:")
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Faqat qiymatlar
	fmt.Println("\nFaqat qiymatlar:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

// demonstrateSliceModification - slice o'zgartirish
func demonstrateSliceModification() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Boshlang'ich:", numbers)

	// Element qo'shish
	numbers = append(numbers, 6)
	fmt.Println("6 qo'shildi:", numbers)

	// Element o'chirish (oxiridan)
	numbers = numbers[:len(numbers)-1]
	fmt.Println("Oxirgi element o'chirildi:", numbers)

	// Element o'chirish (boshidan)
	numbers = numbers[1:]
	fmt.Println("Birinchi element o'chirildi:", numbers)
}

// demonstrateMultiDimensional - ko'p o'lchamli array va slice
func demonstrateMultiDimensional() {
	// 2D array
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println("2D Array:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 2D slice
	slice2D := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("\n2D Slice:")
	for _, row := range slice2D {
		fmt.Println(row)
	}
}

// sumArray - amaliy misol: array yig'indisi
func sumArray(arr [5]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

// findMax - amaliy misol: eng katta son
func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max
}

// findMin - amaliy misol: eng kichik son
func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
	}
	return min
}

// contains - amaliy misol: element borligini tekshirish
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// filterEven - amaliy misol: juft sonlarni filtrlash
func filterEven(numbers []int) []int {
	var result []int
	for _, value := range numbers {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

// reverseSlice - amaliy misol: slice ni teskari qilish
func reverseSlice(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, value := range numbers {
		result[len(numbers)-1-i] = value
	}
	return result
}

// mergeSlices - amaliy misol: ikkita slice ni birlashtirish
func mergeSlices(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}
