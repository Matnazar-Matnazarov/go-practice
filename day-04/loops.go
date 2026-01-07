package main

import "fmt"

// demonstrateForBasic - oddiy for loop
func demonstrateForBasic() {
	fmt.Println("Oddiy for loop (1 dan 5 gacha):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  %d ", i)
	}
	fmt.Println()
}

// demonstrateForCondition - shartli for loop
func demonstrateForCondition() {
	fmt.Println("Shartli for loop:")
	i := 1
	for i <= 5 {
		fmt.Printf("  %d ", i)
		i++
	}
	fmt.Println()
}

// demonstrateForInfinite - cheksiz for loop
func demonstrateForInfinite() {
	fmt.Println("Cheksiz for loop (break bilan):")
	count := 0
	for {
		count++
		if count > 5 {
			break
		}
		fmt.Printf("  %d ", count)
	}
	fmt.Println()
}

// demonstrateForRangeArray - for range array
func demonstrateForRangeArray() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("For range array (index va value):")
	for index, value := range numbers {
		fmt.Printf("  Index %d: %d\n", index, value)
	}

	fmt.Println("\nFor range array (faqat value):")
	for _, value := range numbers {
		fmt.Printf("  %d ", value)
	}
	fmt.Println()

	fmt.Println("\nFor range array (faqat index):")
	for index := range numbers {
		fmt.Printf("  Index %d\n", index)
	}
}

// demonstrateForRangeSlice - for range slice
func demonstrateForRangeSlice() {
	numbers := []int{100, 200, 300, 400, 500}

	fmt.Println("For range slice:")
	for index, value := range numbers {
		fmt.Printf("  Index %d: %d\n", index, value)
	}
}

// demonstrateForRangeString - for range string
func demonstrateForRangeString() {
	text := "Go"

	fmt.Println("For range string (byte):")
	for i := 0; i < len(text); i++ {
		fmt.Printf("  Index %d: %c (byte: %d)\n", i, text[i], text[i])
	}

	fmt.Println("\nFor range string (rune):")
	for index, char := range text {
		fmt.Printf("  Index %d: %c (rune: %d)\n", index, char, char)
	}
}

// demonstrateForRangeMap - for range map
func demonstrateForRangeMap() {
	ages := map[string]int{
		"Matnazar": 25,
		"Ali":      30,
		"Vali":     28,
	}

	fmt.Println("For range map (key va value):")
	for name, age := range ages {
		fmt.Printf("  %s: %d yosh\n", name, age)
	}

	fmt.Println("\nFor range map (faqat key):")
	for name := range ages {
		fmt.Printf("  %s\n", name)
	}
}

// demonstrateNestedLoops - ichki for looplar
func demonstrateNestedLoops() {
	fmt.Println("Ichki for looplar (ko'paytirish jadvali):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("  %d x %d = %d\n", i, j, i*j)
		}
	}
}

// demonstrateBreakContinue - break va continue
func demonstrateBreakContinue() {
	fmt.Println("Break va continue:")

	// Break misoli
	fmt.Println("Break (5 ga yetganda to'xtaydi):")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("  %d ", i)
	}
	fmt.Println()

	// Continue misoli
	fmt.Println("Continue (juft sonlarni o'tkazib yuboradi):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("  %d ", i)
	}
	fmt.Println()
}

// demonstrateLabeledBreak - label bilan break
func demonstrateLabeledBreak() {
	fmt.Println("Label bilan break:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}
}

// demonstrateLabeledContinue - label bilan continue
func demonstrateLabeledContinue() {
	fmt.Println("Label bilan continue:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				continue OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}
}

// sumArrayLoop - amaliy misol: array yig'indisi
func sumArrayLoop(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// sumSliceLoop - amaliy misol: slice yig'indisi
func sumSliceLoop(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// findMaxLoop - amaliy misol: eng katta son
func findMaxLoop(numbers []int) int {
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

// findMinLoop - amaliy misol: eng kichik son
func findMinLoop(numbers []int) int {
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

// countEvenLoop - amaliy misol: juft sonlar soni
func countEvenLoop(numbers []int) int {
	count := 0
	for _, value := range numbers {
		if value%2 == 0 {
			count++
		}
	}
	return count
}

// reverseArrayLoop - amaliy misol: array ni teskari qilish
func reverseArrayLoop(arr [5]int) [5]int {
	result := arr
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// factorialLoop - amaliy misol: faktorial
func factorialLoop(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// fibonacciLoop - amaliy misol: fibonacci ketma-ketligi
func fibonacciLoop(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

// printPattern - amaliy misol: pattern chiqarish
func printPattern(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// searchInSlice - amaliy misol: slice ichida qidirish
func searchInSlice(slice []int, target int) int {
	for index, value := range slice {
		if value == target {
			return index
		}
	}
	return -1
}

// removeDuplicates - amaliy misol: takrorlangan elementlarni olib tashlash
func removeDuplicates(numbers []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, value := range numbers {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}
