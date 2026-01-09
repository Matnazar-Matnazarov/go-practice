package main

import "fmt"

// demonstrateFunctionComposition - funksiyalarni bir-biriga uzatish
func demonstrateFunctionComposition() {
	numbers := []int{1, 2, 3, 4, 5}

	// Funksiyalarni zanjir qilish
	result := multiplyByTwo(addOne(sum(numbers)))
	fmt.Printf("Sonlar: %v\n", numbers)
	fmt.Printf("Yig'indi: %d\n", sum(numbers))
	fmt.Printf("Yig'indi + 1: %d\n", addOne(sum(numbers)))
	fmt.Printf("(Yig'indi + 1) * 2: %d\n", result)
}

// sum - sonlar yig'indisi
func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// addOne - bir qo'shish
func addOne(n int) int {
	return n + 1
}

// multiplyByTwo - ikki ga ko'paytirish
func multiplyByTwo(n int) int {
	return n * 2
}

// demonstrateHigherOrderFunction - higher-order funksiyalar
func demonstrateHigherOrderFunction() {
	numbers := []int{1, 2, 3, 4, 5}

	// Funksiyani parametr sifatida uzatish
	doubled := applyFunction(numbers, multiplyByTwo)
	squared := applyFunction(numbers, square)

	fmt.Printf("Sonlar: %v\n", numbers)
	fmt.Printf("Ikki ga ko'paytirilgan: %v\n", doubled)
	fmt.Printf("Kvadrat: %v\n", squared)
}

// applyFunction - funksiyani har bir elementga qo'llash
func applyFunction(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

// square - kvadrat
func square(n int) int {
	return n * n
}

// demonstrateClosure - closure funksiyalar
func demonstrateClosure() {
	// Counter closure
	counter := createCounter()
	fmt.Println("Counter:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d\n", counter())
	}

	// Multiplier closure
	multiplyBy := createMultiplier(5)
	fmt.Printf("\n5 ga ko'paytirish:\n")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  %d * 5 = %d\n", i, multiplyBy(i))
	}
}

// createCounter - counter closure yaratish
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// createMultiplier - multiplier closure yaratish
func createMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// demonstrateRecursiveFunction - rekursiv funksiyalar
func demonstrateRecursiveFunction() {
	// Faktorial
	fmt.Println("Faktorial (rekursiv):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  %d! = %d\n", i, factorialRecursive(i))
	}

	// Fibonacci
	fmt.Println("\nFibonacci (rekursiv):")
	for i := 0; i < 10; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fibonacciRecursive(i))
	}

	// Yig'indi
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumRecursive(numbers, 0)
	fmt.Printf("\nYig'indi (rekursiv): %v = %d\n", numbers, sum)
}

// factorialRecursive - faktorial (rekursiv)
func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

// fibonacciRecursive - fibonacci (rekursiv)
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// sumRecursive - yig'indi (rekursiv)
func sumRecursive(numbers []int, index int) int {
	if index >= len(numbers) {
		return 0
	}
	return numbers[index] + sumRecursive(numbers, index+1)
}

// demonstrateDeferFunction - defer funksiyalar
func demonstrateDeferFunction() {
	fmt.Println("Defer funksiyalar:")
	defer fmt.Println("  Bu 3-chi chiqadi (defer)")
	fmt.Println("  Bu 1-chi chiqadi")
	defer fmt.Println("  Bu 2-chi chiqadi (defer, teskari tartib)")
}

// demonstrateDeferCleanup - defer bilan tozalash
func demonstrateDeferCleanup() {
	fmt.Println("Defer bilan tozalash:")
	fmt.Println("  Resource ochildi")
	defer fmt.Println("  Resource yopildi (defer)")
	fmt.Println("  Resource ishlatilmoqda...")
}

// demonstrateDeferStack - defer stack (LIFO)
func demonstrateDeferStack() {
	fmt.Println("Defer stack (LIFO):")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  Defer %d\n", i)
	}
	fmt.Println("  Asosiy kod")
}

// demonstrateAnonymousFunction - anonim funksiyalar
func demonstrateAnonymousFunction() {
	// Anonim funksiya
	func() {
		fmt.Println("Anonim funksiya chaqirildi")
	}()

	// Anonim funksiyani o'zgaruvchiga berish
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("Anonim funksiya: 5 + 3 = %d\n", add(5, 3))

	// Anonim funksiya for loop ichida
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Anonim funksiya bilan ko'paytirish:")
	for _, num := range numbers {
		func(n int) {
			fmt.Printf("  %d * 2 = %d\n", n, n*2)
		}(num)
	}
}

// demonstrateFunctionVariadic - variadic funksiyalar (qo'shimcha)
func demonstrateFunctionVariadic() {
	// Variadic funksiya
	fmt.Println("Variadic funksiya:")
	sum1 := variadicSum(1, 2, 3)
	sum2 := variadicSum(1, 2, 3, 4, 5)
	sum3 := variadicSum()

	fmt.Printf("  Sum(1,2,3) = %d\n", sum1)
	fmt.Printf("  Sum(1,2,3,4,5) = %d\n", sum2)
	fmt.Printf("  Sum() = %d\n", sum3)

	// Slice ni variadic funksiyaga uzatish
	numbers := []int{10, 20, 30}
	sum4 := variadicSum(numbers...)
	fmt.Printf("  Sum([10,20,30]...) = %d\n", sum4)
}

// variadicSum - variadic yig'indi
func variadicSum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// demonstrateFunctionAsReturn - funksiyani qaytarish
func demonstrateFunctionAsReturn() {
	// Funksiyani qaytarish
	adder := createAdder(10)
	fmt.Println("Adder funksiyasi (10 qo'shadi):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  %d + 10 = %d\n", i, adder(i))
	}
}

// createAdder - adder funksiyasi yaratish
func createAdder(addend int) func(int) int {
	return func(n int) int {
		return n + addend
	}
}

// demonstrateErrorHandling - error handling (asosiy)
func demonstrateErrorHandling() {
	// Error qaytarish
	result, err := divideSafe(10, 2)
	if err != nil {
		fmt.Printf("Xato: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result2, err2 := divideSafe(10, 0)
	if err2 != nil {
		fmt.Printf("Xato: %v\n", err2)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result2)
	}
}

// divideSafe - xavfsiz bo'lish
func divideSafe(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("nolga bo'lish mumkin emas")
	}
	return a / b, nil
}

// demonstrateFunctionWithStruct - struct bilan funksiyalar
func demonstrateFunctionWithStruct() {
	person := Person{"Matnazar", 25, "Toshkent"}

	// Funksiya struct qabul qiladi
	greetPerson(person)

	// Funksiya struct qaytaradi
	newPerson := updatePersonAge(person, 26)
	fmt.Printf("Yangi yosh: %d\n", newPerson.Age)
}

// greetPerson - person ni salomlash
func greetPerson(p Person) {
	fmt.Printf("Salom, %s! (%d yosh, %s)\n", p.Name, p.Age, p.City)
}

// updatePersonAge - person yoshini yangilash
func updatePersonAge(p Person, newAge int) Person {
	p.Age = newAge
	return p
}
