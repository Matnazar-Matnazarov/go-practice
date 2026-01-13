package main

import "fmt"

// Counter struct
type Counter struct {
	value int
}

// Increment - Value receiver
func (c Counter) Increment() {
	c.value++ // Nusxa o'zgaradi
}

// IncrementPtr - Pointer receiver
func (c *Counter) IncrementPtr() {
	c.value++ // Asl o'zgaradi
}

// demonstratePointerBasic - Pointer asoslari
func demonstratePointerBasic() {
	fmt.Println("Pointer asoslari:")

	var x int = 42
	var p *int = &x // Pointer yaratish

	fmt.Printf("  x qiymati: %d\n", x)
	fmt.Printf("  x manzili: %p\n", &x)
	fmt.Printf("  p qiymati (manzil): %p\n", p)
	fmt.Printf("  p orqali qiymat (*p): %d\n", *p) // Dereference
}

// demonstratePointerModify - Pointer orqali o'zgartirish
func demonstratePointerModify() {
	fmt.Println("Pointer orqali o'zgartirish:")

	x := 10
	fmt.Printf("  O'zgartirishdan oldin: %d\n", x)

	p := &x
	*p = 20 // Pointer orqali o'zgartirish
	fmt.Printf("  O'zgartirishdan keyin: %d\n", x)
}

// demonstratePointerNil - Nil pointer
func demonstratePointerNil() {
	fmt.Println("Nil pointer:")

	var p *int
	fmt.Printf("  p == nil: %t\n", p == nil)

	// *p = 10 // Bu panic qiladi!
	// fmt.Printf("  *p: %d\n", *p) // Bu ham panic qiladi!
}

// demonstratePointerNew - new() funksiyasi
func demonstratePointerNew() {
	fmt.Println("new() funksiyasi:")

	p := new(int) // Yangi int yaratadi va pointer qaytaradi
	fmt.Printf("  p qiymati: %p\n", p)
	fmt.Printf("  *p qiymati: %d\n", *p) // 0 (zero value)

	*p = 100
	fmt.Printf("  *p yangi qiymati: %d\n", *p)
}

// demonstratePointerFunction - Funksiyaga pointer uzatish
func demonstratePointerFunction() {
	fmt.Println("Funksiyaga pointer uzatish:")

	x := 10
	fmt.Printf("  Funksiyadan oldin: %d\n", x)

	modifyValue(&x) // Pointer uzatish
	fmt.Printf("  Funksiyadan keyin: %d\n", x)
}

// modifyValue - Pointer qabul qiluvchi funksiya
func modifyValue(p *int) {
	*p = 999
}

// demonstratePointerReturn - Pointer qaytarish
func demonstratePointerReturn() {
	fmt.Println("Pointer qaytarish:")

	p := createInt(42)
	fmt.Printf("  Qaytarilgan pointer: %p\n", p)
	fmt.Printf("  Qiymat: %d\n", *p)
}

// createInt - Yangi int yaratadi va pointer qaytaradi
func createInt(value int) *int {
	return &value
}

// demonstratePointerStruct - Struct pointer
func demonstratePointerStruct() {
	fmt.Println("Struct pointer:")

	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "Ali", Age: 25}
	fmt.Printf("  p1: %+v\n", p1)

	p2 := &p1
	p2.Age = 30 // Pointer orqali o'zgartirish
	fmt.Printf("  p1 o'zgardi: %+v\n", p1)
	fmt.Printf("  p2: %+v\n", *p2)
}

// demonstratePointerArray - Array pointer
func demonstratePointerArray() {
	fmt.Println("Array pointer:")

	arr := [3]int{1, 2, 3}
	p := &arr

	fmt.Printf("  arr: %v\n", arr)
	fmt.Printf("  p[0]: %d\n", p[0]) // Pointer orqali elementga kirish
	p[1] = 999                       // Pointer orqali o'zgartirish
	fmt.Printf("  arr o'zgardi: %v\n", arr)
}

// demonstratePointerSlice - Slice pointer (odatda kerak emas)
func demonstratePointerSlice() {
	fmt.Println("Slice pointer (odatda kerak emas):")

	slice := []int{1, 2, 3}
	p := &slice

	fmt.Printf("  slice: %v\n", slice)
	fmt.Printf("  (*p)[0]: %d\n", (*p)[0]) // Qavslar kerak
	(*p)[1] = 999
	fmt.Printf("  slice o'zgardi: %v\n", slice)
}

// demonstratePointerComparison - Pointer taqqoslash
func demonstratePointerComparison() {
	fmt.Println("Pointer taqqoslash:")

	x := 10
	p1 := &x
	p2 := &x

	fmt.Printf("  p1 == p2: %t\n", p1 == p2) // True - bir xil manzil

	y := 10
	p3 := &y
	fmt.Printf("  p1 == p3: %t\n", p1 == p3)     // False - turli manzillar
	fmt.Printf("  *p1 == *p3: %t\n", *p1 == *p3) // True - bir xil qiymat
}

// demonstratePointerMethod - Pointer receiver
func demonstratePointerMethod() {
	fmt.Println("Pointer receiver:")

	// Value receiver
	c1 := Counter{value: 10}
	c1.Increment()                                 // Nusxa o'zgaradi
	fmt.Printf("  Value receiver: %d\n", c1.value) // 10 (o'zgarmadi)

	// Pointer receiver
	c2 := Counter{value: 10}
	c2.IncrementPtr()                                // Asl o'zgaradi
	fmt.Printf("  Pointer receiver: %d\n", c2.value) // 11 (o'zgardi)
}

// demonstratePointerPractical - Amaliy misol: Swap
func demonstratePointerPractical() {
	fmt.Println("Amaliy misol: Swap funksiyasi:")

	a, b := 10, 20
	fmt.Printf("  Swap dan oldin: a=%d, b=%d\n", a, b)

	swap(&a, &b)
	fmt.Printf("  Swap dan keyin: a=%d, b=%d\n", a, b)
}

// swap - Ikkita qiymatni almashtirish
func swap(x, y *int) {
	*x, *y = *y, *x
}
