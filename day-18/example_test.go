package main

import "fmt"

// ExampleAdd - Add funksiyasi uchun misol
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

// ExampleSubtract - Subtract funksiyasi uchun misol
func ExampleSubtract() {
	result := Subtract(10, 3)
	fmt.Println(result)
	// Output: 7
}

// ExampleMultiply - Multiply funksiyasi uchun misol
func ExampleMultiply() {
	result := Multiply(3, 4)
	fmt.Println(result)
	// Output: 12
}

// ExampleMax - Max funksiyasi uchun misol
func ExampleMax() {
	result := Max(5, 10)
	fmt.Println(result)
	// Output: 10
}

// ExampleMin - Min funksiyasi uchun misol
func ExampleMin() {
	result := Min(5, 10)
	fmt.Println(result)
	// Output: 5
}

// ExampleIsEven - IsEven funksiyasi uchun misol
func ExampleIsEven() {
	fmt.Println(IsEven(2))
	fmt.Println(IsEven(3))
	// Output:
	// true
	// false
}

// ExampleFactorial - Factorial funksiyasi uchun misol
func ExampleFactorial() {
	fmt.Println(Factorial(5))
	// Output: 120
}

// ExampleFibonacci - Fibonacci funksiyasi uchun misol
func ExampleFibonacci() {
	fmt.Println(Fibonacci(10))
	// Output: 55
}
