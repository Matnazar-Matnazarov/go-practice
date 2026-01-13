package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Circle metodlari
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle metodlari
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Reader interface
type Reader interface {
	Read() string
}

// Writer interface
type Writer interface {
	Write(string)
}

// ReadWriter interface (composition)
type ReadWriter interface {
	Reader
	Writer
}

// File struct
type File struct {
	content string
}

// File metodlari
func (f *File) Read() string {
	return f.content
}

func (f *File) Write(s string) {
	f.content = s
}

// Animal interface
type Animal interface {
	Sound() string
}

// Dog struct
type Dog struct{}

// Dog metodlari
func (d Dog) Sound() string {
	return "Woof!"
}

// Cat struct
type Cat struct{}

// Cat metodlari
func (c Cat) Sound() string {
	return "Meow!"
}

// Person struct (Stringer uchun)
type Person struct {
	Name string
	Age  int
}

// Person String() metodi
func (p Person) String() string {
	return fmt.Sprintf("%s (%d yosh)", p.Name, p.Age)
}

// CustomError struct (Error uchun)
type CustomError struct {
	Code    int
	Message string
}

// CustomError Error() metodi
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// ReadWriterWithClose interface (embedding uchun)
type ReadWriterWithClose interface {
	Reader
	Writer
	Close()
}

// FileWithClose struct
type FileWithClose struct {
	content string
	closed  bool
}

// FileWithClose metodlari
func (f *FileWithClose) Read() string {
	if f.closed {
		return ""
	}
	return f.content
}

func (f *FileWithClose) Write(s string) {
	if !f.closed {
		f.content = s
	}
}

func (f *FileWithClose) Close() {
	f.closed = true
}

// demonstrateInterfaceBasic - Interface asoslari
func demonstrateInterfaceBasic() {
	fmt.Println("Interface asoslari:")

	var s Shape
	s = Circle{Radius: 5}
	fmt.Printf("  Circle area: %.2f\n", s.Area())
	fmt.Printf("  Circle perimeter: %.2f\n", s.Perimeter())

	s = Rectangle{Width: 4, Height: 6}
	fmt.Printf("  Rectangle area: %.2f\n", s.Area())
	fmt.Printf("  Rectangle perimeter: %.2f\n", s.Perimeter())
}

// demonstrateInterfaceEmpty - Empty interface
func demonstrateInterfaceEmpty() {
	fmt.Println("Empty interface (interface{}):")

	var i interface{}
	i = 42
	fmt.Printf("  int: %v\n", i)

	i = "hello"
	fmt.Printf("  string: %v\n", i)

	i = true
	fmt.Printf("  bool: %v\n", i)
}

// demonstrateInterfaceTypeAssertion - Type assertion
func demonstrateInterfaceTypeAssertion() {
	fmt.Println("Type assertion:")

	var i interface{} = "hello"

	// Type assertion
	s, ok := i.(string)
	if ok {
		fmt.Printf("  String: %s\n", s)
	}

	// Type assertion (panic qilishi mumkin)
	// s2 := i.(string) // OK
	// n := i.(int)     // Panic!

	// Type switch
	switch v := i.(type) {
	case string:
		fmt.Printf("  Type: string, Value: %s\n", v)
	case int:
		fmt.Printf("  Type: int, Value: %d\n", v)
	default:
		fmt.Printf("  Type: unknown, Value: %v\n", v)
	}
}

// demonstrateInterfaceComposition - Interface composition
func demonstrateInterfaceComposition() {
	fmt.Println("Interface composition:")

	var rw ReadWriter = &File{content: "Hello"}
	fmt.Printf("  Read: %s\n", rw.Read())
	rw.Write("World")
	fmt.Printf("  Read after Write: %s\n", rw.Read())
}

// demonstrateInterfacePolymorphism - Polymorphism
func demonstrateInterfacePolymorphism() {
	fmt.Println("Polymorphism:")

	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Printf("  %s\n", animal.Sound())
	}
}

// demonstrateInterfaceNil - Nil interface
func demonstrateInterfaceNil() {
	fmt.Println("Nil interface:")

	var i interface{}
	fmt.Printf("  i == nil: %t\n", i == nil)

	var s *string
	i = s
	fmt.Printf("  i == nil: %t\n", i == nil) // False! (nil pointer, lekin interface nil emas)
}

// demonstrateInterfacePractical - Amaliy misol: Sort
func demonstrateInterfacePractical() {
	fmt.Println("Amaliy misol: Sort interface:")

	people := []Person{
		{Name: "Ali", Age: 25},
		{Name: "Vali", Age: 30},
		{Name: "Hasan", Age: 20},
	}

	fmt.Println("  Sort dan oldin:")
	for _, p := range people {
		fmt.Printf("    %s: %d\n", p.Name, p.Age)
	}

	// Oddiy sort (bubble sort)
	for i := 0; i < len(people)-1; i++ {
		for j := 0; j < len(people)-i-1; j++ {
			if people[j].Age > people[j+1].Age {
				people[j], people[j+1] = people[j+1], people[j]
			}
		}
	}

	fmt.Println("  Sort dan keyin:")
	for _, p := range people {
		fmt.Printf("    %s: %d\n", p.Name, p.Age)
	}
}

// demonstrateInterfaceStringer - Stringer interface
func demonstrateInterfaceStringer() {
	fmt.Println("Stringer interface:")

	p := Person{Name: "Ali", Age: 25}
	fmt.Printf("  %s\n", p) // String() avtomatik chaqiriladi
	fmt.Printf("  %v\n", p) // String() avtomatik chaqiriladi
}

// demonstrateInterfaceError - Error interface
func demonstrateInterfaceError() {
	fmt.Println("Error interface:")

	var err error = &CustomError{Code: 404, Message: "Not Found"}
	fmt.Printf("  Error: %v\n", err)
}

// demonstrateInterfaceEmbedding - Interface embedding
func demonstrateInterfaceEmbedding() {
	fmt.Println("Interface embedding:")

	var rw ReadWriterWithClose = &FileWithClose{content: "Hello"}
	fmt.Printf("  Read: %s\n", rw.Read())
	rw.Write("World")
	fmt.Printf("  Read after Write: %s\n", rw.Read())
	rw.Close()
	fmt.Printf("  Read after Close: %s\n", rw.Read())
}
