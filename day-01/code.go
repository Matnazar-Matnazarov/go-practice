package main

import "fmt"

// main funksiyasi - har bir Go dasturining boshlanish nuqtasi
// Bu funksiya dastur ishga tushganda avtomatik chaqiriladi
func main() {

	// 1. SALOM DUNYO - Eng oddiy Go dasturi

	fmt.Println("=== Salom, Go dasturlash tili! ===")
	fmt.Println()

	// 2. O'ZGARUVCHILAR (VARIABLES)

	// Usul 1: var bilan to'liq e'lon qilish (tip ko'rsatiladi)
	var ism string = "Matnazar"
	var yosh int = 25
	var shahar string = "Toshkent"

	fmt.Println("--- O'zgaruvchilar ---")
	fmt.Println("Ism:", ism)
	fmt.Println("Yosh:", yosh)
	fmt.Println("Shahar:", shahar)
	fmt.Println()

	// Usul 2: var bilan qisqa e'lon (tip avtomatik aniqlanadi)
	var telefon = "+998901234567"
	var email = "matnazar@example.com"

	fmt.Println("Telefon:", telefon)
	fmt.Println("Email:", email)
	fmt.Println()

	// Usul 3: Qisqa sintaksis (:=) - faqat funksiya ichida ishlaydi
	// Bu eng qulay usul
	kasb := "Dasturchi"
	maosh := 5000.50
	ishlaydi := true

	fmt.Println("Kasb:", kasb)
	fmt.Println("Maosh:", maosh)
	fmt.Println("Ishlaydi:", ishlaydi)
	fmt.Println()

	// 3. MA'LUMOT TURLARI (DATA TYPES)

	fmt.Println("--- Ma'lumot turlari ---")

	// String (matn)
	var matn string = "Bu matn turi"
	fmt.Printf("String: %s\n", matn)

	// Integer (butun son)
	var butunSon int = 42
	fmt.Printf("Integer: %d\n", butunSon)

	// Float (o'nlik son)
	var onlikSon float64 = 3.14159
	fmt.Printf("Float: %.2f\n", onlikSon)

	// Boolean (mantiqiy)
	var rost bool = true
	var yolgon bool = false
	fmt.Printf("Boolean (rost): %t\n", rost)
	fmt.Printf("Boolean (yolgon): %t\n", yolgon)
	fmt.Println()

	// 4. PRINT FUNKSIYALARI

	fmt.Println("--- Print funksiyalari ---")

	// fmt.Print() - oddiy chiqarish, yangi qator qo'shmaydi
	fmt.Print("Bu birinchi matn. ")
	fmt.Print("Bu ikkinchi matn. ")
	fmt.Print("Yangi qator yo'q.\n")

	// fmt.Println() - chiqarish va yangi qator qo'shadi
	fmt.Println("Bu yangi qator bilan chiqadi")
	fmt.Println("Bu ham yangi qator bilan")

	// fmt.Printf() - formatlangan chiqarish
	ism2 := "Go"
	versiya := 1.21
	fmt.Printf("Dasturlash tili: %s, Versiya: %.2f\n", ism2, versiya)
	fmt.Println()

	// 5. O'ZGARUVCHILARNI QAYTA ISHLATISH

	fmt.Println("--- O'zgaruvchilarni yangilash ---")

	hisob := 100
	fmt.Printf("Boshlang'ich hisob: %d\n", hisob)

	hisob = 150 // Yangi qiymat berish (= ishlatiladi, := emas)
	fmt.Printf("Yangi hisob: %d\n", hisob)

	hisob = hisob + 50 // Matematik amallar
	fmt.Printf("Yana qo'shildi: %d\n", hisob)
	fmt.Println()

	// 6. KONSTANTALAR (CONSTANTS)

	fmt.Println("--- Konstanta ---")

	// const - o'zgarmas qiymatlar
	const pi float64 = 3.14159
	const davlat string = "O'zbekiston"

	fmt.Printf("Pi soni: %.5f\n", pi)
	fmt.Printf("Davlat: %s\n", davlat)

	// const qiymatini o'zgartirib bo'lmaydi
	// pi = 3.14 // Bu xato bo'ladi!
	fmt.Println()

	// 7. MATEMATIK AMALLAR

	fmt.Println("--- Matematik amallar ---")

	a := 10
	b := 5

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // % - qoldiq
	fmt.Println()

	// 8. MATN BILAN ISHLASH

	fmt.Println("--- Matn bilan ishlash ---")

	familiya := "Matnazarov"
	ism3 := "Matnazar"

	// Matnlarni birlashtirish (+)
	toliqIsm := ism3 + " " + familiya
	fmt.Println("To'liq ism:", toliqIsm)

	// fmt.Sprintf() - formatlangan matn yaratish
	matn2 := fmt.Sprintf("Ism: %s, Familiya: %s", ism3, familiya)
	fmt.Println("Formatlangan matn:", matn2)
	fmt.Println()

	// 9. YAKUNIY XULOSA

	fmt.Println("=== Kun 1 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Package va import")
	fmt.Println("  ✓ O'zgaruvchilar (var, :=)")
	fmt.Println("  ✓ Ma'lumot turlari (string, int, float64, bool)")
	fmt.Println("  ✓ Print funksiyalari")
	fmt.Println("  ✓ Konstanta")
	fmt.Println("  ✓ Matematik amallar")
}
