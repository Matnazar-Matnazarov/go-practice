package main

import "fmt"

// Funksiya turlari (Function types)

// Operation - Funksiya turi
type Operation func(int, int) int

// Predicate - Funksiya turi (shart tekshirish)
type Predicate func(int) bool

// Transformer - Funksiya turi (transformatsiya)
type Transformer func(int) int

// Callback - Funksiya turi (callback)
type Callback func(string)

// Processor interface - Higher-order funksiyalar uchun
type Processor interface {
	Process(int) int
}

// FuncProcessor - Funksiya interface ni implement qiladi
type FuncProcessor func(int) int

// Process - FuncProcessor metod
func (f FuncProcessor) Process(n int) int {
	return f(n)
}

// demonstrateFunctionTypes - Funksiya turlari
func demonstrateFunctionTypes() {
	fmt.Println("Funksiya turlari (Function types):")

	// Funksiya turi e'lon qilish
	var add Operation = func(a, b int) int {
		return a + b
	}

	var multiply Operation = func(a, b int) int {
		return a * b
	}

	fmt.Printf("  add(5, 3) = %d\n", add(5, 3))
	fmt.Printf("  multiply(5, 3) = %d\n", multiply(5, 3))

	// Funksiya turini parametr sifatida uzatish
	result := calculate(10, 5, add)
	fmt.Printf("  calculate(10, 5, add) = %d\n", result)

	result = calculate(10, 5, multiply)
	fmt.Printf("  calculate(10, 5, multiply) = %d\n", result)
}

// calculate - Funksiya turini parametr sifatida qabul qiladi
func calculate(a, b int, op Operation) int {
	return op(a, b)
}

// demonstrateFunctionValues - Funksiya qiymatlari
func demonstrateFunctionValues() {
	fmt.Println("Funksiya qiymatlari (Function values):")

	// Funksiyani o'zgaruvchiga saqlash
	add := func(a, b int) int {
		return a + b
	}

	fmt.Printf("  add(10, 20) = %d\n", add(10, 20))

	// Funksiyani boshqa o'zgaruvchiga nusxalash
	addCopy := add
	fmt.Printf("  addCopy(10, 20) = %d\n", addCopy(10, 20))

	// Funksiyani slice ga saqlash
	operations := []Operation{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a - b },
		func(a, b int) int { return a * b },
	}

	for i, op := range operations {
		fmt.Printf("  operations[%d](10, 5) = %d\n", i, op(10, 5))
	}
}

// demonstrateClosureWithPointer - Closure bilan pointer
func demonstrateClosureWithPointer() {
	fmt.Println("Closure bilan pointer:")

	// Closure - tashqi o'zgaruvchilarga kirish
	counter := 0
	increment := func() {
		counter++
	}

	increment()
	increment()
	fmt.Printf("  counter = %d\n", counter)

	// Closure bilan pointer
	value := 10
	modify := func() {
		value = 20 // Tashqi o'zgaruvchini o'zgartirish
	}

	modify()
	fmt.Printf("  value = %d\n", value)

	// Closure - bir nechta funksiyalar
	x := 0
	inc := func() { x++ }
	dec := func() { x-- }
	get := func() int { return x }

	inc()
	inc()
	dec()
	fmt.Printf("  get() = %d\n", get())
}

// demonstrateHigherOrderWithInterface - Higher-order funksiyalar bilan interface
func demonstrateHigherOrderWithInterface() {
	fmt.Println("Higher-order funksiyalar bilan interface:")

	// Processor yaratish
	double := FuncProcessor(func(n int) int {
		return n * 2
	})

	square := FuncProcessor(func(n int) int {
		return n * n
	})

	fmt.Printf("  double.Process(5) = %d\n", double.Process(5))
	fmt.Printf("  square.Process(5) = %d\n", square.Process(5))

	// Higher-order funksiya
	process := func(p Processor, n int) int {
		return p.Process(n)
	}

	fmt.Printf("  process(double, 5) = %d\n", process(double, 5))
	fmt.Printf("  process(square, 5) = %d\n", process(square, 5))
}

// demonstrateCallbackFunctions - Callback funksiyalar
func demonstrateCallbackFunctions() {
	fmt.Println("Callback funksiyalar:")

	// Callback funksiya
	process := func(name string, callback Callback) {
		fmt.Printf("  Processing: %s\n", name)
		callback("Done")
	}

	// Callback funksiyani uzatish
	process("Task 1", func(msg string) {
		fmt.Printf("  Callback: %s\n", msg)
	})

	// Callback funksiyani o'zgaruvchiga saqlash
	onComplete := func(msg string) {
		fmt.Printf("  OnComplete: %s\n", msg)
	}

	process("Task 2", onComplete)
}

// demonstrateFunctionAsParameter - Funksiya parametr sifatida
func demonstrateFunctionAsParameter() {
	fmt.Println("Funksiya parametr sifatida:")

	numbers := []int{1, 2, 3, 4, 5}

	// Map - har bir elementni transformatsiya qilish
	doubled := mapFunc(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("  mapFunc(numbers, double): %v\n", doubled)

	// Filter - shartga mos elementlarni qoldirish
	evens := filterFunc(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("  filterFunc(numbers, even): %v\n", evens)

	// Reduce - barcha elementlarni birlashtirish
	sum := reduceFunc(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("  reduceFunc(numbers, sum): %d\n", sum)
}

// mapFunc - Map funksiyasi
func mapFunc(slice []int, fn Transformer) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// filterFunc - Filter funksiyasi
func filterFunc(slice []int, fn Predicate) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// reduceFunc - Reduce funksiyasi
func reduceFunc(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// demonstrateFunctionReturn - Funksiya qaytarish
func demonstrateFunctionReturn() {
	fmt.Println("Funksiya qaytarish:")

	// Funksiya qaytarish
	createAdder := func(n int) func(int) int {
		return func(x int) int {
			return x + n
		}
	}

	add5 := createAdder(5)
	add10 := createAdder(10)

	fmt.Printf("  add5(10) = %d\n", add5(10))
	fmt.Printf("  add10(10) = %d\n", add10(10))

	// Funksiya qaytarish - multiplier
	createMultiplier := func(factor int) Transformer {
		return func(n int) int {
			return n * factor
		}
	}

	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Printf("  double(5) = %d\n", double(5))
	fmt.Printf("  triple(5) = %d\n", triple(5))
}
