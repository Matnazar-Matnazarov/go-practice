package main

import "fmt"

// demonstrateMapBasic - map asoslari
func demonstrateMapBasic() {
	// Map e'lon qilish - make() bilan
	ages1 := make(map[string]int)
	ages1["Matnazar"] = 25
	fmt.Println("Make() bilan:", ages1)

	// Map literal sintaksis
	ages := map[string]int{
		"Matnazar": 25,
		"Ali":      30,
		"Vali":     28,
	}

	fmt.Println("Map:", ages)
	fmt.Printf("Map uzunligi: %d\n", len(ages))
}

// demonstrateMapOperations - map operatsiyalari
func demonstrateMapOperations() {
	// Map yaratish
	student := make(map[string]string)
	student["name"] = "Matnazar"
	student["age"] = "25"
	student["city"] = "Toshkent"

	fmt.Println("Student map:", student)

	// Elementlarga kirish
	fmt.Printf("Ism: %s\n", student["name"])
	fmt.Printf("Yosh: %s\n", student["age"])

	// Elementni o'zgartirish
	student["city"] = "Samarqand"
	fmt.Println("O'zgartirilgan map:", student)

	// Element qo'shish
	student["email"] = "matnazar@example.com"
	fmt.Println("Yangi element qo'shilgan:", student)

	// Element o'chirish
	delete(student, "age")
	fmt.Println("Element o'chirilgan:", student)
}

// demonstrateMapCheck - map element borligini tekshirish
func demonstrateMapCheck() {
	ages := map[string]int{
		"Matnazar": 25,
		"Ali":      30,
	}

	// Element borligini tekshirish
	age, exists := ages["Matnazar"]
	if exists {
		fmt.Printf("Matnazar yoshi: %d\n", age)
	} else {
		fmt.Println("Matnazar topilmadi")
	}

	// Element yo'qligini tekshirish
	age2, exists2 := ages["Vali"]
	if exists2 {
		fmt.Printf("Vali yoshi: %d\n", age2)
	} else {
		fmt.Println("Vali topilmadi")
	}

	// Qisqa sintaksis
	if age3, ok := ages["Ali"]; ok {
		fmt.Printf("Ali yoshi: %d\n", age3)
	}
}

// demonstrateMapIteration - map iteratsiyasi
func demonstrateMapIteration() {
	students := map[string]int{
		"Matnazar": 95,
		"Ali":      85,
		"Vali":     90,
		"Bobur":    88,
	}

	// For range loop
	fmt.Println("Barcha talabalar:")
	for name, score := range students {
		fmt.Printf("  %s: %d ball\n", name, score)
	}

	// Faqat kalitlar
	fmt.Println("\nFaqat ismlar:")
	for name := range students {
		fmt.Printf("  %s\n", name)
	}

	// Faqat qiymatlar
	fmt.Println("\nFaqat ballar:")
	for _, score := range students {
		fmt.Printf("  %d\n", score)
	}
}

// demonstrateMapTypes - turli tipdagi maplar
func demonstrateMapTypes() {
	// String to int
	stringIntMap := map[string]int{
		"bir":  1,
		"ikki": 2,
		"uch":  3,
	}
	fmt.Println("String to int:", stringIntMap)

	// Int to string
	intStringMap := map[int]string{
		1: "bir",
		2: "ikki",
		3: "uch",
	}
	fmt.Println("Int to string:", intStringMap)

	// String to slice
	stringSliceMap := map[string][]int{
		"juft":   {2, 4, 6, 8},
		"toq":    {1, 3, 5, 7},
		"birlik": {1, 11, 21},
	}
	fmt.Println("String to slice:", stringSliceMap)

	// Nested map
	nestedMap := map[string]map[string]int{
		"O'zbekiston": {
			"Toshkent":  3000000,
			"Samarqand": 500000,
		},
		"Qozog'iston": {
			"Olmaota":    2000000,
			"Qarag'anda": 500000,
		},
	}
	fmt.Println("Nested map:", nestedMap)
}

// demonstrateMapCopy - map nusxalash
func demonstrateMapCopy() {
	original := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Nusxalash
	copied := make(map[string]int)
	for key, value := range original {
		copied[key] = value
	}

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// O'zgartirish
	copied["d"] = 4
	fmt.Println("O'zgartirilgan copied:", copied)
	fmt.Println("Original o'zgarmadi:", original)
}

// demonstrateMapClear - map ni tozalash
func demonstrateMapClear() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println("Boshlang'ich map:", myMap)
	fmt.Printf("Uzunlik: %d\n", len(myMap))

	// Tozalash
	for key := range myMap {
		delete(myMap, key)
	}

	fmt.Println("Tozalangan map:", myMap)
	fmt.Printf("Uzunlik: %d\n", len(myMap))
}

// countWords - amaliy misol: so'zlar sonini hisoblash
func countWords(text string) map[string]int {
	words := make(map[string]int)
	word := ""

	for _, char := range text {
		if char == ' ' || char == ',' || char == '.' {
			if word != "" {
				words[word]++
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words[word]++
	}

	return words
}

// findMaxValue - amaliy misol: eng katta qiymat
func findMaxValue(m map[string]int) (string, int) {
	if len(m) == 0 {
		return "", 0
	}

	maxKey := ""
	maxValue := 0
	first := true

	for key, value := range m {
		if first || value > maxValue {
			maxValue = value
			maxKey = key
			first = false
		}
	}

	return maxKey, maxValue
}

// mergeMaps - amaliy misol: ikkita map ni birlashtirish
func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	// Birinchi map
	for key, value := range map1 {
		result[key] = value
	}

	// Ikkinchi map
	for key, value := range map2 {
		result[key] = value
	}

	return result
}

// filterMap - amaliy misol: map ni filtrlash
func filterMap(m map[string]int, minValue int) map[string]int {
	result := make(map[string]int)

	for key, value := range m {
		if value >= minValue {
			result[key] = value
		}
	}

	return result
}

// getKeys - amaliy misol: barcha kalitlarni olish
func getKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// getValues - amaliy misol: barcha qiymatlarni olish
func getValues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// invertMap - amaliy misol: map ni teskari qilish
func invertMap(m map[string]int) map[int]string {
	result := make(map[int]string)
	for key, value := range m {
		result[value] = key
	}
	return result
}
