package main

import "fmt"

// Counter struct - metodlar uchun
type Counter struct {
	value int
}

// Increment - Value receiver
func (c Counter) Increment() {
	c.value++
}

// IncrementPtr - Pointer receiver
func (c *Counter) IncrementPtr() {
	c.value++
}

// GetValue - Value receiver
func (c Counter) GetValue() int {
	return c.value
}

// SetValue - Pointer receiver
func (c *Counter) SetValue(v int) {
	c.value = v
}

// Add - Pointer receiver
func (c *Counter) Add(n int) *Counter {
	c.value += n
	return c // Method chaining uchun
}

// Multiply - Pointer receiver
func (c *Counter) Multiply(n int) *Counter {
	c.value *= n
	return c // Method chaining uchun
}

// Person struct - method sets uchun
type Person struct {
	Name string
	Age  int
}

// GetName - Value receiver
func (p Person) GetName() string {
	return p.Name
}

// SetName - Pointer receiver
// Eslatma: Value receiver bilan SetName ishlamaydi
func (p *Person) SetName(name string) {
	p.Name = name
}

// GetAge - Value receiver
func (p Person) GetAge() int {
	return p.Age
}

// SetAge - Pointer receiver
func (p *Person) SetAge(age int) {
	p.Age = age
}

// String - Value receiver (Stringer interface)
func (p Person) String() string {
	return fmt.Sprintf("%s (%d yosh)", p.Name, p.Age)
}

// Calculator struct - method chaining uchun
type Calculator struct {
	result float64
}

// Add - Method chaining
func (c *Calculator) Add(n float64) *Calculator {
	c.result += n
	return c
}

// Subtract - Method chaining
func (c *Calculator) Subtract(n float64) *Calculator {
	c.result -= n
	return c
}

// Multiply - Method chaining
func (c *Calculator) Multiply(n float64) *Calculator {
	c.result *= n
	return c
}

// Divide - Method chaining
func (c *Calculator) Divide(n float64) *Calculator {
	if n != 0 {
		c.result /= n
	}
	return c
}

// GetResult - Result olish
func (c *Calculator) GetResult() float64 {
	return c.result
}

// Reset - Reset qilish
func (c *Calculator) Reset() *Calculator {
	c.result = 0
	return c
}

// demonstrateMethodSets - Method sets
func demonstrateMethodSets() {
	fmt.Println("Method sets qoidalari:")

	// Value type method set:
	// - Barcha value receiver metodlar
	// - Barcha pointer receiver metodlar (Go avtomatik qo'shadi)

	// Pointer type method set:
	// - Barcha value receiver metodlar
	// - Barcha pointer receiver metodlar

	var p1 Person = Person{Name: "Ali", Age: 25}
	var p2 *Person = &Person{Name: "Vali", Age: 30}

	// Value type - value receiver metodlar ishlaydi
	fmt.Printf("  p1.GetName(): %s\n", p1.GetName())

	// Value type - pointer receiver metodlar ham ishlaydi (Go avtomatik)
	p1.SetName("Aliyev")
	fmt.Printf("  p1.SetName() dan keyin: %s\n", p1.GetName())

	// Pointer type - value receiver metodlar ishlaydi
	fmt.Printf("  p2.GetName(): %s\n", p2.GetName())

	// Pointer type - pointer receiver metodlar ishlaydi
	p2.SetName("Valiyev")
	fmt.Printf("  p2.SetName() dan keyin: %s\n", p2.GetName())
}

// demonstratePointerReceiverRules - Pointer receiver qoidalari
func demonstratePointerReceiverRules() {
	fmt.Println("Pointer receiver qoidalari:")

	c1 := Counter{value: 10}
	c2 := &Counter{value: 20}

	// Value receiver - nusxa o'zgaradi
	c1.Increment()
	fmt.Printf("  c1.Increment() dan keyin: %d (o'zgarmadi)\n", c1.GetValue())

	// Pointer receiver - asl o'zgaradi
	c1.IncrementPtr()
	fmt.Printf("  c1.IncrementPtr() dan keyin: %d (o'zgardi)\n", c1.GetValue())

	// Pointer type - pointer receiver
	c2.IncrementPtr()
	fmt.Printf("  c2.IncrementPtr() dan keyin: %d (o'zgardi)\n", c2.GetValue())

	// Qoida:
	// - O'zgartirish kerak bo'lsa - Pointer receiver
	// - Faqat o'qish kerak bo'lsa - Value receiver
	// - Struct katta bo'lsa - Pointer receiver (samaradorlik)
}

// demonstrateValueVsPointerReceiver - Value vs Pointer receiver
func demonstrateValueVsPointerReceiver() {
	fmt.Println("Value vs Pointer receiver:")

	// Value receiver - nusxa yaratadi
	p1 := Person{Name: "Ali", Age: 25}
	p1.SetName("Aliyev") // Go avtomatik &p1 qiladi
	fmt.Printf("  Value receiver bilan SetName: %s\n", p1.GetName())

	// Pointer receiver - to'g'ridan-to'g'ri o'zgartirish
	p2 := &Person{Name: "Vali", Age: 30}
	p2.SetName("Valiyev")
	fmt.Printf("  Pointer receiver bilan SetName: %s\n", p2.GetName())

	// Eslatma: Go avtomatik konvertatsiya qiladi
	// p1.SetName() = (&p1).SetName()
	// p2.GetName() = (*p2).GetName()
}

// demonstrateMethodExpressions - Method expressions
func demonstrateMethodExpressions() {
	fmt.Println("Method expressions:")

	p := Person{Name: "Ali", Age: 25}

	// Method expression - funksiya sifatida
	// Type.MethodName - funksiya qaytaradi
	getNameFunc := Person.GetName
	fmt.Printf("  getNameFunc(p): %s\n", getNameFunc(p))

	setNameFunc := (*Person).SetName
	setNameFunc(&p, "Aliyev")
	fmt.Printf("  setNameFunc(&p, \"Aliyev\") dan keyin: %s\n", p.GetName())

	// Pointer type method expression
	getNamePtrFunc := (*Person).GetName
	fmt.Printf("  getNamePtrFunc(&p): %s\n", getNamePtrFunc(&p))
}

// demonstrateMethodValues - Method values
func demonstrateMethodValues() {
	fmt.Println("Method values:")

	p := Person{Name: "Ali", Age: 25}

	// Method value - o'zgaruvchiga bog'langan metod
	getName := p.GetName
	fmt.Printf("  getName(): %s\n", getName())

	// Method value - pointer bilan
	p2 := &Person{Name: "Vali", Age: 30}
	setName := p2.SetName
	setName("Valiyev")
	fmt.Printf("  setName(\"Valiyev\") dan keyin: %s\n", p2.GetName())

	// Method value - closure kabi ishlaydi
	counter := Counter{value: 10}
	increment := counter.IncrementPtr
	increment()
	fmt.Printf("  increment() dan keyin: %d\n", counter.GetValue())
}

// demonstrateMethodChaining - Method chaining
func demonstrateMethodChaining() {
	fmt.Println("Method chaining:")

	calc := &Calculator{result: 0}
	result := calc.Add(10).Multiply(2).Subtract(5).Divide(3).GetResult()
	fmt.Printf("  calc.Add(10).Multiply(2).Subtract(5).Divide(3) = %.2f\n", result)

	// Counter bilan method chaining
	counter := &Counter{value: 5}
	counter.Add(10).Multiply(2)
	fmt.Printf("  counter.Add(10).Multiply(2) = %d\n", counter.GetValue())
}

// demonstrateInterfaceMethods - Interface metodlar
func demonstrateInterfaceMethods() {
	fmt.Println("Interface metodlar:")

	// Interface metodlari - faqat metodlar nomi va signature
	type Stringer interface {
		String() string
	}

	var s Stringer = Person{Name: "Ali", Age: 25}
	fmt.Printf("  Stringer.String(): %s\n", s.String())

	// Interface - barcha metodlar implement qilinishi kerak
	type Getter interface {
		GetName() string
		GetAge() int
	}

	var g Getter = Person{Name: "Vali", Age: 30}
	fmt.Printf("  Getter.GetName(): %s\n", g.GetName())
	fmt.Printf("  Getter.GetAge(): %d\n", g.GetAge())
}
