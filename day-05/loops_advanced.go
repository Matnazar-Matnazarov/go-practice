package main

import "fmt"

// demonstrateLoopOptimization - loop optimizatsiyasi
func demonstrateLoopOptimization() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Oddiy loop
	fmt.Println("Oddiy loop:")
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Printf("  Yig'indi: %d\n", sum)

	// Range loop (optimallashtirilgan)
	fmt.Println("\nRange loop:")
	sum2 := 0
	for _, num := range numbers {
		sum2 += num
	}
	fmt.Printf("  Yig'indi: %d\n", sum2)
}

// demonstrateLoopWithCondition - shartli loop
func demonstrateLoopWithCondition() {
	fmt.Println("Shartli loop (juft sonlar):")
	i := 1
	for i <= 20 {
		if i%2 == 0 {
			fmt.Printf("  %d ", i)
		}
		i++
	}
	fmt.Println()
}

// demonstrateLoopEarlyExit - erta chiqish
func demonstrateLoopEarlyExit() {
	fmt.Println("Erta chiqish (10 dan katta son topilganda):")
	numbers := []int{1, 3, 5, 12, 7, 9, 11}

	for _, num := range numbers {
		if num > 10 {
			fmt.Printf("  %d topildi, loop to'xtatildi\n", num)
			break
		}
		fmt.Printf("  %d ", num)
	}
	fmt.Println()
}

// demonstrateLoopSkip - elementlarni o'tkazib yuborish
func demonstrateLoopSkip() {
	fmt.Println("Elementlarni o'tkazib yuborish (3 ga bo'linadigan sonlar):")
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Printf("  %d ", i)
	}
	fmt.Println()
}

// demonstrateLoopNested - ichki looplar (qo'shimcha misollar)
func demonstrateLoopNested() {
	fmt.Println("Ichki looplar (ko'paytirish jadvali 5x5):")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("  %d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}

// demonstrateLoopWithSlice - slice bilan loop
func demonstrateLoopWithSlice() {
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Slice bilan loop (index va value):")
	for index, value := range numbers {
		fmt.Printf("  Index %d: %d\n", index, value)
	}

	fmt.Println("\nSlice elementlarini o'zgartirish:")
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Printf("  O'zgartirilgan: %v\n", numbers)
}

// demonstrateLoopWithMap - map bilan loop
func demonstrateLoopWithMap() {
	ages := map[string]int{
		"Matnazar": 25,
		"Ali":      30,
		"Vali":     28,
		"Bobur":    22,
	}

	fmt.Println("Map bilan loop:")
	for name, age := range ages {
		fmt.Printf("  %s: %d yosh\n", name, age)
	}

	fmt.Println("\nMap elementlarini o'zgartirish:")
	for name := range ages {
		ages[name]++
	}
	fmt.Printf("  O'zgartirilgan: %v\n", ages)
}

// demonstrateLoopWithString - string bilan loop
func demonstrateLoopWithString() {
	text := "Go"

	fmt.Println("String bilan loop (byte):")
	for i := 0; i < len(text); i++ {
		fmt.Printf("  Index %d: %c (byte: %d)\n", i, text[i], text[i])
	}

	fmt.Println("\nString bilan loop (rune):")
	for index, char := range text {
		fmt.Printf("  Index %d: %c (rune: %d)\n", index, char, char)
	}
}

// demonstrateLoopPattern - pattern chiqarish
func demonstrateLoopPattern() {
	fmt.Println("Pattern 1 (uchburchak):")
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println("\nPattern 2 (teskari uchburchak):")
	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println("\nPattern 3 (piramida):")
	for i := 1; i <= rows; i++ {
		// Bo'sh joylar
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Yulduzlar
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// demonstrateLoopSearch - qidirish algoritmlari
func demonstrateLoopSearch() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70}

	// Linear search
	target := 40
	index := linearSearch(numbers, target)
	if index != -1 {
		fmt.Printf("Linear search: %d topildi, index: %d\n", target, index)
	} else {
		fmt.Printf("Linear search: %d topilmadi\n", target)
	}

	// Binary search (tartiblangan array uchun)
	sorted := []int{10, 20, 30, 40, 50, 60, 70}
	index2 := binarySearch(sorted, target)
	if index2 != -1 {
		fmt.Printf("Binary search: %d topildi, index: %d\n", target, index2)
	} else {
		fmt.Printf("Binary search: %d topilmadi\n", target)
	}
}

// linearSearch - linear qidirish
func linearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

// binarySearch - binary qidirish
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// demonstrateLoopSort - sort algoritmlari
func demonstrateLoopSort() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Boshlang'ich: %v\n", numbers)

	// Bubble sort
	sorted := bubbleSort(numbers)
	fmt.Printf("Bubble sort: %v\n", sorted)
}

// bubbleSort - bubble sort algoritmi
func bubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// demonstrateLoopWithStruct - struct bilan loop
func demonstrateLoopWithStruct() {
	people := []Person{
		{"Matnazar", 25, "Toshkent"},
		{"Ali", 30, "Samarqand"},
		{"Vali", 28, "Buxoro"},
	}

	fmt.Println("Struct bilan loop:")
	for i, person := range people {
		fmt.Printf("  %d. %s, %d yosh, %s\n", i+1, person.Name, person.Age, person.City)
	}

	// Struct fieldlarni o'zgartirish
	fmt.Println("\nYoshlarni yangilash:")
	for i := range people {
		people[i].Age++
	}
	for _, person := range people {
		fmt.Printf("  %s: %d yosh\n", person.Name, person.Age)
	}
}

// demonstrateLoopPerformance - loop performance
func demonstrateLoopPerformance() {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// Pre-allocate slice
	result := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	fmt.Printf("Pre-allocated slice: %d element\n", len(result))
}

// demonstrateLoopWithError - loop bilan error handling
func demonstrateLoopWithError() {
	numbers := []float64{10, 5, 0, 20, 15}

	fmt.Println("Loop bilan error handling:")
	for i, num := range numbers {
		result, err := divideSafe(num, 2)
		if err != nil {
			fmt.Printf("  Index %d: Xato - %v\n", i, err)
		} else {
			fmt.Printf("  Index %d: %.2f / 2 = %.2f\n", i, num, result)
		}
	}
}

// demonstrateLoopTransformation - loop bilan transformatsiya
func demonstrateLoopTransformation() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map transformatsiya
	doubled := make([]int, len(numbers))
	for i, num := range numbers {
		doubled[i] = num * 2
	}
	fmt.Printf("Ikki ga ko'paytirilgan: %v\n", doubled)

	// Filter transformatsiya
	evens := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Printf("Juft sonlar: %v\n", evens)

	// Reduce transformatsiya
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Yig'indi: %d\n", sum)
}
