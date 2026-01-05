package main

import (
	"fmt"
	"strings"
)

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	// 1. Oddiy funksiya - parametr va qaytarish qiymati yo'q
	fmt.Println("Kun 2: Funksiyalar va String operatsiyalari")
	fmt.Println()

	// Funksiyani chaqirish
	greet()
	fmt.Println()

	// 2. Parametrli funksiya - funksiyaga ma'lumot uzatish
	fmt.Println("Parametrli funksiyalar:")

	// Bitta parametrli funksiya
	greetPerson("Matnazar")
	greetPerson("Go")

	// Ko'p parametrli funksiya
	printInfo("Matnazar", 25, "Toshkent")
	fmt.Println()

	// 3. Qaytarish qiymatli funksiya - funksiyadan natija olish
	fmt.Println("Qaytarish qiymatli funksiyalar:")

	// Oddiy qaytarish qiymati
	sumResult := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sumResult)

	// Funksiya natijasini to'g'ridan-to'g'ri ishlatish
	fmt.Printf("15 * 3 = %d\n", multiply(15, 3))
	fmt.Printf("100 / 4 = %.2f\n", divide(100, 4))
	fmt.Println()

	// 4. Ko'p qaytarish qiymati - bir nechta natija qaytarish
	fmt.Println("Ko'p qaytarish qiymati:")

	// Bo'lish va qoldiq
	quotient, remainder := divideAndRemainder(17, 5)
	fmt.Printf("17 / 5 = %d, qoldiq: %d\n", quotient, remainder)

	// To'rtburchak perimetri va yuzasi
	perimeter, area := rectangle(10, 5)
	fmt.Printf("To'rtburchak (10x5): Perimeter=%d, Yuzasi=%d\n", perimeter, area)

	// Blank identifier (_) - keraksiz qiymatni e'tiborsiz qoldirish
	_, onlyRemainder := divideAndRemainder(20, 7)
	fmt.Printf("20 / 7, faqat qoldiq: %d\n", onlyRemainder)
	fmt.Println()

	// 5. Nomlangan qaytarish qiymatlar - qaytarish qiymatlariga nom berish
	fmt.Println("Nomlangan qaytarish qiymatlar:")

	sum2, difference := calculate(50, 30)
	fmt.Printf("50 va 30: Yig'indi=%d, Ayirma=%d\n", sum2, difference)

	product, division := mathOperations(12, 4)
	fmt.Printf("12 va 4: Ko'paytma=%d, Bo'linma=%d\n", product, division)
	fmt.Println()

	// 6. Variadic funksiya - cheksiz parametr qabul qilish
	fmt.Println("Variadic funksiyalar:")

	// Bir nechta son yig'indisi
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", total)

	total2 := sum(10, 20, 30)
	fmt.Printf("10+20+30 = %d\n", total2)

	// Bir nechta matn birlashtirish
	text := concatenate("Salom", " ", "Go", " ", "dasturlash", " ", "tili")
	fmt.Println("Birlashtirilgan matn:", text)
	fmt.Println()

	// 7. Funksiyalarni bir-biriga uzatish - funksiya ichida funksiya chaqirish
	fmt.Println("Funksiyalarni bir-biriga uzatish:")

	result := sumOfSquares(3, 4)
	fmt.Printf("3² + 4² = %d\n", result)

	// Murakkab hisob-kitob
	answer := complexCalculation(10, 5)
	fmt.Printf("(10+5)² - (10-5)² = %d\n", answer)
	fmt.Println()

	// 8. Funksiya tiplari - funksiyani o'zgaruvchi sifatida ishlatish
	fmt.Println("Funksiya tiplari:")

	// Funksiyani o'zgaruvchiga berish
	calcFunction := add
	fmt.Printf("Funksiya orqali: 15 + 25 = %d\n", calcFunction(15, 25))

	// Boshqa funksiyaga uzatish
	result2 := calculateWithFunction(8, 2, multiply)
	fmt.Printf("Funksiya parametr orqali: 8 * 2 = %d\n", result2)
	fmt.Println()

	// 9. String operatsiyalari - keng qamrovli matn bilan ishlash
	fmt.Println("String operatsiyalari:")

	// String uzunligi
	text1 := "Go dasturlash tili"
	length := len(text1)
	fmt.Printf("Matn: '%s', Uzunligi: %d\n", text1, length)

	// String birlashtirish
	text2 := "Salom"
	text3 := "Dunyo"
	combined := text2 + " " + text3
	fmt.Printf("Birlashtirilgan: %s\n", combined)

	// strings.Contains() - matn ichida qidirish
	contains := strings.Contains(text1, "Go")
	fmt.Printf("'%s' ichida 'Go' bor: %t\n", text1, contains)

	// strings.HasPrefix() - boshlanishini tekshirish
	hasPrefix := strings.HasPrefix(text1, "Go")
	fmt.Printf("'%s' 'Go' bilan boshlanadi: %t\n", text1, hasPrefix)

	// strings.HasSuffix() - oxirini tekshirish
	hasSuffix := strings.HasSuffix(text1, "tili")
	fmt.Printf("'%s' 'tili' bilan tugaydi: %t\n", text1, hasSuffix)

	// strings.ToUpper() - katta harflarga o'tkazish
	upper := strings.ToUpper(text1)
	fmt.Printf("Katta harflar: %s\n", upper)

	// strings.ToLower() - kichik harflarga o'tkazish
	lower := strings.ToLower("GO DASTURLASH TILI")
	fmt.Printf("Kichik harflar: %s\n", lower)

	// strings.TrimSpace() - bo'sh joylarni olib tashlash
	trimmed := strings.TrimSpace("  Go dasturlash tili  ")
	fmt.Printf("Trim qilingan: '%s'\n", trimmed)

	// strings.Trim() - belgilarni olib tashlash
	trimmed2 := strings.Trim("!!!Go!!!", "!")
	fmt.Printf("Trim qilingan: '%s'\n", trimmed2)

	// strings.Split() - matnni bo'lish
	words := strings.Split("Go Python Java", " ")
	fmt.Printf("Bo'lingan so'zlar: %v\n", words)

	// strings.Join() - matnlarni birlashtirish
	joined := strings.Join([]string{"Go", "Python", "Java"}, " | ")
	fmt.Printf("Birlashtirilgan: %s\n", joined)

	// strings.Replace() - matnni almashtirish
	replaced := strings.Replace(text1, "Go", "Golang", 1)
	fmt.Printf("Almashtirilgan: %s\n", replaced)

	// strings.ReplaceAll() - barcha o'rinlarni almashtirish
	replacedAll := strings.ReplaceAll("Go Go Go", "Go", "Golang")
	fmt.Printf("Barcha almashtirilgan: %s\n", replacedAll)

	// strings.Index() - belgi pozitsiyasini topish
	index := strings.Index(text1, "dasturlash")
	fmt.Printf("'dasturlash' pozitsiyasi: %d\n", index)

	// strings.Count() - belgi sonini hisoblash
	count := strings.Count(text1, "o")
	fmt.Printf("'o' harfi soni: %d\n", count)

	// strings.Repeat() - matnni takrorlash
	repeated := strings.Repeat("Go ", 3)
	fmt.Printf("Takrorlangan: %s\n", repeated)

	// strings.Compare() - matnlarni solishtirish
	compare1 := strings.Compare("Go", "Python")
	compare2 := strings.Compare("Go", "Go") // nolint:gocritic // misol uchun
	compare3 := strings.Compare("Python", "Go")
	fmt.Printf("Compare('Go', 'Python'): %d\n", compare1)
	fmt.Printf("Compare('Go', 'Go'): %d\n", compare2)
	fmt.Printf("Compare('Python', 'Go'): %d\n", compare3)

	// strings.EqualFold() - katta-kichik harfni e'tiborsiz solishtirish
	equal := strings.EqualFold("Go", "GO")
	fmt.Printf("EqualFold('Go', 'GO'): %t\n", equal)

	// strings.Fields() - bo'sh joylar bo'yicha bo'lish
	fields := strings.Fields("Go Python Java Rust")
	fmt.Printf("Fields: %v\n", fields)

	// strings.Title() - har bir so'zning birinchi harfini katta qilish
	title := strings.Title("go dasturlash tili")
	fmt.Printf("Title: %s\n", title)
	fmt.Println()

	// 10. Amaliy misollar - haqiqiy hayotdan misollar
	fmt.Println("Amaliy misollar:")

	// Ism va familiyani birlashtirish
	fullName := combineNames("Matnazar", "Matnazarov")
	fmt.Println("To'liq ism:", fullName)

	// Yoshi hisoblash
	age := calculateAge(1999)
	fmt.Printf("1999-yilda tug'ilgan, hozirgi yosh: %d\n", age)

	// Doira yuzasi va perimetri
	circleArea, circlePerimeter := circle(5.0)
	fmt.Printf("Doira (r=5): Yuzasi=%.2f, Perimetri=%.2f\n", circleArea, circlePerimeter)

	// Uchburchak yuzasi (Geron formulasi)
	triangleArea := triangleArea(3, 4, 5)
	fmt.Printf("Uchburchak (3,4,5): Yuzasi=%.2f\n", triangleArea)

	// Email tekshirish
	email := "user@example.com"
	isValidEmail := validateEmail(email)
	fmt.Printf("Email '%s' to'g'ri: %t\n", email, isValidEmail)

	// Matnni formatlash
	formatted := formatText("go dasturlash tili")
	fmt.Printf("Formatlangan matn: %s\n", formatted)
	fmt.Println()

	// 11. Yakuniy xulosa
	fmt.Println("Kun 2 yakunlandi!")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Oddiy funksiyalar")
	fmt.Println("  ✓ Parametrli funksiyalar")
	fmt.Println("  ✓ Qaytarish qiymatli funksiyalar")
	fmt.Println("  ✓ Ko'p qaytarish qiymatlari")
	fmt.Println("  ✓ Nomlangan qaytarish qiymatlar")
	fmt.Println("  ✓ Variadic funksiyalar")
	fmt.Println("  ✓ Funksiya tiplari")
	fmt.Println("  ✓ String operatsiyalari (len, Contains, Split, Join, Replace, va boshqalar)")
}
