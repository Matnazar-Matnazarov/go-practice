package main

import "fmt"


// 1. GENERICS ASOSLARI


// Print - any type uchun print funksiyasi
func Print[T any](value T) {
	fmt.Printf("  Value: %v (type: %T)\n", value, value)
}

// demonstrateTypeParameters - Type parameters
func demonstrateTypeParameters() {
	Print(42)
	Print("hello")
	Print(3.14)
	Print(true)
}

// Find - comparable type lar uchun find funksiyasi
func Find[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// demonstrateTypeConstraints - Type constraints
func demonstrateTypeConstraints() {
	numbers := []int{1, 2, 3, 4, 5}
	index := Find(numbers, 3)
	fmt.Printf("  Find(3) in [1,2,3,4,5]: index = %d\n", index)

	strings := []string{"apple", "banana", "cherry"}
	index2 := Find(strings, "banana")
	fmt.Printf("  Find(\"banana\") in [\"apple\",\"banana\",\"cherry\"]: index = %d\n", index2)
}

// Swap - multiple type parameters
func Swap[T, U any](a T, b U) (U, T) {
	return b, a
}

// demonstrateMultipleTypeParameters - Multiple type parameters
func demonstrateMultipleTypeParameters() {
	x, y := Swap(10, "hello")
	fmt.Printf("  Swap(10, \"hello\"): %v, %v\n", x, y)

	a, b := Swap("world", 42)
	fmt.Printf("  Swap(\"world\", 42): %v, %v\n", a, b)
}

// Stack - generic stack struct
type Stack[T any] struct {
	items []T
}

// Push - stack ga element qo'shish
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop - stack dan element olish
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// demonstrateTypeParametersWithStruct - Type parameters bilan struct
func demonstrateTypeParametersWithStruct() {
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Printf("  Pop: %d\n", intStack.Pop())

	stringStack := &Stack[string]{}
	stringStack.Push("a")
	stringStack.Push("b")
	fmt.Printf("  Pop: %s\n", stringStack.Pop())
}


// 2. TYPE CONSTRAINTS (KENGAYTIRILGAN)


// Number - number type lar uchun constraint
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Add - Number constraint bilan
func Add[T Number](a, b T) T {
	return a + b
}

// demonstrateInterfaceConstraints - Interface constraints
func demonstrateInterfaceConstraints() {
	sum1 := Add(10, 20)
	fmt.Printf("  Add(10, 20) = %d\n", sum1)

	sum2 := Add(3.14, 2.86)
	fmt.Printf("  Add(3.14, 2.86) = %.2f\n", sum2)
}

// Stringer - method constraint
type Stringer interface {
	String() string
}

// Person - Stringer interface ni implement qiladi
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d yosh)", p.Name, p.Age)
}

// PrintStringer - Stringer constraint bilan
func PrintStringer[T Stringer](value T) {
	fmt.Printf("  %s\n", value.String())
}

// demonstrateMethodConstraints - Method constraints
func demonstrateMethodConstraints() {
	person := Person{Name: "Ali", Age: 25}
	PrintStringer(person)
}

// Addable - custom constraint
type Addable interface {
	int | float64 | string
}

// Sum - Addable constraint bilan
func Sum[T Addable](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// demonstrateCustomConstraints - Custom constraints
func demonstrateCustomConstraints() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Printf("  Sum([1,2,3,4,5]) = %d\n", sum)

	floats := []float64{1.1, 2.2, 3.3}
	sum2 := Sum(floats)
	fmt.Printf("  Sum([1.1,2.2,3.3]) = %.2f\n", sum2)
}

// MyInt - custom int type
type MyInt int

// Integer - ~ operator bilan constraint
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Double - Integer constraint bilan
func Double[T Integer](value T) T {
	return value * 2
}

// demonstrateTildeOperator - ~ Operator
func demonstrateTildeOperator() {
	var x MyInt = 10
	result := Double(x)
	fmt.Printf("  Double(MyInt(10)) = %d\n", result)
}


// 3. GENERICS BILAN FUNKSIYALAR


// Map - generic map funksiyasi
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// demonstrateGenericFunctions - Generic funksiyalar
func demonstrateGenericFunctions() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("  Map([1,2,3,4,5], *2) = %v\n", doubled)

	strings := []string{"hello", "world"}
	lengths := Map(strings, func(s string) int {
		return len(s)
	})
	fmt.Printf("  Map([\"hello\",\"world\"], len) = %v\n", lengths)
}

// Filter - generic filter funksiyasi
func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// demonstrateGenericFilter - Generic filter
func demonstrateGenericFilter() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("  Filter([1..10], even) = %v\n", evens)

	strings := []string{"apple", "banana", "cherry", "date"}
	long := Filter(strings, func(s string) bool {
		return len(s) > 5
	})
	fmt.Printf("  Filter([\"apple\",\"banana\",\"cherry\",\"date\"], len>5) = %v\n", long)
}

// Reduce - generic reduce funksiyasi
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// demonstrateGenericReduce - Generic reduce
func demonstrateGenericReduce() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("  Reduce([1,2,3,4,5], sum) = %d\n", sum)

	product := Reduce(numbers, 1, func(acc, n int) int {
		return acc * n
	})
	fmt.Printf("  Reduce([1,2,3,4,5], product) = %d\n", product)
}

// FindGeneric - generic find funksiyasi
func FindGeneric[T comparable](slice []T, value T) (T, bool) {
	for _, v := range slice {
		if v == value {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// demonstrateGenericFind - Generic find
func demonstrateGenericFind() {
	numbers := []int{1, 2, 3, 4, 5}
	value, found := FindGeneric(numbers, 3)
	if found {
		fmt.Printf("  FindGeneric([1,2,3,4,5], 3) = %d, found: true\n", value)
	}

	value2, found2 := FindGeneric(numbers, 10)
	if !found2 {
		fmt.Printf("  FindGeneric([1,2,3,4,5], 10) = %d, found: false\n", value2)
	}
}


// 4. GENERICS BILAN DATA STRUCTURES


// NewStack - yangi stack yaratish
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// demonstrateGenericStack - Generic stack
func demonstrateGenericStack() {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Printf("  Push(1), Push(2), Push(3)\n")
	fmt.Printf("  Pop: %d\n", stack.Pop())
	fmt.Printf("  Pop: %d\n", stack.Pop())
	fmt.Printf("  Pop: %d\n", stack.Pop())
}

// Queue - generic queue struct
type Queue[T any] struct {
	items []T
}

// NewQueue - yangi queue yaratish
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue - queue ga element qo'shish
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue - queue dan element olish
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// demonstrateGenericQueue - Generic queue
func demonstrateGenericQueue() {
	queue := NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	fmt.Printf("  Enqueue(\"first\"), Enqueue(\"second\"), Enqueue(\"third\")\n")
	item1, _ := queue.Dequeue()
	fmt.Printf("  Dequeue: %s\n", item1)
	item2, _ := queue.Dequeue()
	fmt.Printf("  Dequeue: %s\n", item2)
	item3, _ := queue.Dequeue()
	fmt.Printf("  Dequeue: %s\n", item3)
}

// GenericMap - generic map struct
type GenericMap[K comparable, V any] struct {
	items map[K]V
}

// NewGenericMap - yangi generic map yaratish
func NewGenericMap[K comparable, V any]() *GenericMap[K, V] {
	return &GenericMap[K, V]{items: make(map[K]V)}
}

// Set - map ga key-value qo'shish
func (m *GenericMap[K, V]) Set(key K, value V) {
	m.items[key] = value
}

// Get - map dan value olish
func (m *GenericMap[K, V]) Get(key K) (V, bool) {
	value, ok := m.items[key]
	return value, ok
}

// demonstrateGenericMap - Generic map
func demonstrateGenericMap() {
	myMap := NewGenericMap[string, int]()
	myMap.Set("one", 1)
	myMap.Set("two", 2)
	myMap.Set("three", 3)

	value1, _ := myMap.Get("one")
	fmt.Printf("  Get(\"one\") = %d\n", value1)

	value2, _ := myMap.Get("two")
	fmt.Printf("  Get(\"two\") = %d\n", value2)
}


// 5. GENERICS BILAN INTERFACE


// Container - generic interface
type Container[T any] interface {
	Add(T)
	Get() (T, bool)
	Size() int
}

// StackContainer - Container interface ni implement qiladi
func (s *Stack[T]) Add(item T) {
	s.Push(item)
}

// Get - Stack dan element olish
func (s *Stack[T]) Get() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Size - Stack o'lchami
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// demonstrateGenericInterface - Generic interface
func demonstrateGenericInterface() {
	var container Container[int] = NewStack[int]()
	container.Add(1)
	container.Add(2)
	container.Add(3)
	fmt.Printf("  Container size: %d\n", container.Size())
	value, _ := container.Get()
	fmt.Printf("  Container.Get() = %d\n", value)
}

// Process - type assertion bilan generics
func Process[T any](value T) {
	switch v := any(value).(type) {
	case int:
		fmt.Printf("  Integer: %d\n", v)
	case string:
		fmt.Printf("  String: %s\n", v)
	case float64:
		fmt.Printf("  Float: %.2f\n", v)
	default:
		fmt.Printf("  Unknown: %v\n", v)
	}
}

// demonstrateTypeAssertionWithGenerics - Type assertion bilan generics
func demonstrateTypeAssertionWithGenerics() {
	Process(42)
	Process("hello")
	Process(3.14)
}


// 6. MUHIM QOIDALAR


// demonstrateTypeInference - Type inference
func demonstrateTypeInference() {
	// Type inference
	Print(42)      // T = int
	Print("hello") // T = string

	// Explicit type
	Print[int](42)
	Print[string]("hello")
}

// GetZero - zero value qaytarish
func GetZero[T any]() T {
	var zero T
	return zero
}

// demonstrateZeroValue - Zero value
func demonstrateZeroValue() {
	zeroInt := GetZero[int]()
	fmt.Printf("  GetZero[int]() = %d\n", zeroInt)

	zeroString := GetZero[string]()
	fmt.Printf("  GetZero[string]() = \"%s\"\n", zeroString)
}

// Modify - pointer types bilan
func Modify[T any](value *T, newValue T) {
	*value = newValue
}

// demonstratePointerTypes - Pointer types
func demonstratePointerTypes() {
	x := 10
	fmt.Printf("  Before Modify: x = %d\n", x)
	Modify(&x, 20)
	fmt.Printf("  After Modify(&x, 20): x = %d\n", x)
}
