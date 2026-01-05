package main

import "fmt"

// main funksiyasi - har bir Go dasturining boshlanish nuqtasi
// Bu funksiya dastur ishga tushganda avtomatik chaqiriladi
func main() {
	// 1. Salom dunyo - eng oddiy Go dasturi
	fmt.Println("Salom, Go dasturlash tili!")
	fmt.Println()

	// 2. O'zgaruvchilar (Variables)
	// Usul 1: var bilan to'liq e'lon qilish (tip ko'rsatiladi)
	var name string = "Matnazar"
	var age int = 25
	var city string = "Toshkent"

	fmt.Println("O'zgaruvchilar:")
	fmt.Println("Ism:", name)
	fmt.Println("Yosh:", age)
	fmt.Println("Shahar:", city)
	fmt.Println()

	// Usul 2: var bilan qisqa e'lon (tip avtomatik aniqlanadi)
	var phone = "+998901234567"
	var email = "matnazar@example.com"

	fmt.Println("Telefon:", phone)
	fmt.Println("Email:", email)
	fmt.Println()

	// Usul 3: Qisqa sintaksis (:=) - faqat funksiya ichida ishlaydi
	// Bu eng qulay usul
	profession := "Dasturchi"
	salary := 5000.50
	isWorking := true

	fmt.Println("Kasb:", profession)
	fmt.Println("Maosh:", salary)
	fmt.Println("Ishlaydi:", isWorking)
	fmt.Println()

	// 3. Ma'lumot turlari (Data Types)
	fmt.Println("Ma'lumot turlari:")

	// String (matn)
	var text string = "Bu matn turi"
	fmt.Printf("String: %s\n", text)

	// Integer (butun son)
	var integer int = 42
	fmt.Printf("Integer: %d\n", integer)

	// Float (o'nlik son)
	var floatNum float64 = 3.14159
	fmt.Printf("Float: %.2f\n", floatNum)

	// Boolean (mantiqiy)
	var trueValue bool = true
	var falseValue bool = false
	fmt.Printf("Boolean (rost): %t\n", trueValue)
	fmt.Printf("Boolean (yolgon): %t\n", falseValue)
	fmt.Println()

	// 4. Print funksiyalari
	fmt.Println("Print funksiyalari:")

	// fmt.Print() - oddiy chiqarish, yangi qator qo'shmaydi
	fmt.Print("Bu birinchi matn. ")
	fmt.Print("Bu ikkinchi matn. ")
	fmt.Print("Yangi qator yo'q.\n")

	// fmt.Println() - chiqarish va yangi qator qo'shadi
	fmt.Println("Bu yangi qator bilan chiqadi")
	fmt.Println("Bu ham yangi qator bilan")

	// fmt.Printf() - formatlangan chiqarish
	language := "Go"
	version := 1.21
	fmt.Printf("Dasturlash tili: %s, Versiya: %.2f\n", language, version)
	fmt.Println()

	// 5. O'zgaruvchilarni qayta ishlatish
	fmt.Println("O'zgaruvchilarni yangilash:")

	balance := 100
	fmt.Printf("Boshlang'ich hisob: %d\n", balance)

	balance = 150 // Yangi qiymat berish (= ishlatiladi, := emas)
	fmt.Printf("Yangi hisob: %d\n", balance)

	balance += 50 // Matematik amallar
	fmt.Printf("Yana qo'shildi: %d\n", balance)
	fmt.Println()

	// 6. Konstanta (Constants)
	fmt.Println("Konstanta:")

	// const - o'zgarmas qiymatlar
	const pi float64 = 3.14159
	const country string = "O'zbekiston"

	fmt.Printf("Pi soni: %.5f\n", pi)
	fmt.Printf("Davlat: %s\n", country)

	// const qiymatini o'zgartirib bo'lmaydi
	// pi = 3.14 // Bu xato bo'ladi!
	fmt.Println()

	// 7. Matematik amallar
	fmt.Println("Matematik amallar:")

	a := 10
	b := 5

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // % - qoldiq
	fmt.Println()

	// 8. Matn bilan ishlash (String operations)
	fmt.Println("Matn bilan ishlash:")

	lastName := "Matnazarov"
	firstName := "Matnazar"

	// Matnlarni birlashtirish (+)
	fullName := firstName + " " + lastName
	fmt.Println("To'liq ism:", fullName)

	// fmt.Sprintf() - formatlangan matn yaratish
	formattedText := fmt.Sprintf("Ism: %s, Familiya: %s", firstName, lastName)
	fmt.Println("Formatlangan matn:", formattedText)
	fmt.Println()

	// 9. Yakuniy xulosa
	fmt.Println("Kun 1 yakunlandi!")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Package va import")
	fmt.Println("  ✓ O'zgaruvchilar (var, :=)")
	fmt.Println("  ✓ Ma'lumot turlari (string, int, float64, bool)")
	fmt.Println("  ✓ Print funksiyalari")
	fmt.Println("  ✓ Konstanta")
	fmt.Println("  ✓ Matematik amallar")
	fmt.Println("  ✓ Matn bilan ishlash")
}
