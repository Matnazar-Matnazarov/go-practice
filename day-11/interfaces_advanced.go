package main

import "fmt"

// Reader interface
type Reader interface {
	Read() string
}

// Writer interface
type Writer interface {
	Write(string)
}

// Closer interface
type Closer interface {
	Close()
}

// ReadWriter interface (composition)
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriteCloser interface (composition)
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// File struct - interface implementatsiyasi
type File struct {
	content string
	closed  bool
}

// File metodlari
func (f *File) Read() string {
	if f.closed {
		return ""
	}
	return f.content
}

func (f *File) Write(s string) {
	if !f.closed {
		f.content = s
	}
}

func (f *File) Close() {
	f.closed = true
}

// Buffer struct - interface implementatsiyasi
type Buffer struct {
	data []byte
}

func (b *Buffer) Read() string {
	return string(b.data)
}

func (b *Buffer) Write(s string) {
	b.data = []byte(s)
}

// Adder struct
type Adder struct {
	value int
}

func (a *Adder) Process(n int) int {
	return n + a.value
}

// Multiplier struct
type Multiplier struct {
	factor int
}

func (m *Multiplier) Process(n int) int {
	return n * m.factor
}

// demonstrateInterfaceImplementation - Interface implementatsiyasi
func demonstrateInterfaceImplementation() {
	fmt.Println("Interface implementatsiyasi:")

	// File - ReadWriter interface ni implement qiladi
	var rw ReadWriter = &File{content: "Hello"}
	fmt.Printf("  Read: %s\n", rw.Read())
	rw.Write("World")
	fmt.Printf("  Read after Write: %s\n", rw.Read())

	// Buffer - ReadWriter interface ni implement qiladi
	var rw2 ReadWriter = &Buffer{data: []byte("Test")}
	fmt.Printf("  Buffer Read: %s\n", rw2.Read())
	rw2.Write("New Content")
	fmt.Printf("  Buffer Read after Write: %s\n", rw2.Read())

	// ReadWriteCloser
	var rwc ReadWriteCloser = &File{content: "Initial"}
	fmt.Printf("  RWC Read: %s\n", rwc.Read())
	rwc.Write("Updated")
	fmt.Printf("  RWC Read after Write: %s\n", rwc.Read())
	rwc.Close()
	fmt.Printf("  RWC Read after Close: %s\n", rwc.Read())
}

// demonstrateTypeAssertionWithMethods - Type assertion bilan metodlar
func demonstrateTypeAssertionWithMethods() {
	fmt.Println("Type assertion bilan metodlar:")

	var p Processor

	// Adder
	p = &Adder{value: 10}
	fmt.Printf("  Adder.Process(5) = %d\n", p.Process(5))

	// Type assertion
	if adder, ok := p.(*Adder); ok {
		fmt.Printf("  Type: Adder, value: %d\n", adder.value)
	}

	// Multiplier
	p = &Multiplier{factor: 3}
	fmt.Printf("  Multiplier.Process(5) = %d\n", p.Process(5))

	// Type assertion
	if multiplier, ok := p.(*Multiplier); ok {
		fmt.Printf("  Type: Multiplier, factor: %d\n", multiplier.factor)
	}

	// Type switch
	switch v := p.(type) {
	case *Adder:
		fmt.Printf("  Type switch: Adder with value %d\n", v.value)
	case *Multiplier:
		fmt.Printf("  Type switch: Multiplier with factor %d\n", v.factor)
	default:
		fmt.Printf("  Type switch: Unknown type\n")
	}
}

// demonstrateInterfaceConversion - Interface conversion
func demonstrateInterfaceConversion() {
	fmt.Println("Interface conversion:")

	// File - ReadWriter
	file := &File{content: "Hello"}
	var rw ReadWriter = file

	// ReadWriter dan Reader ga
	var r Reader = rw
	fmt.Printf("  Reader.Read(): %s\n", r.Read())

	// ReadWriter dan Writer ga
	var w Writer = rw
	w.Write("World")
	fmt.Printf("  Writer.Write() dan keyin Read(): %s\n", r.Read())

	// ReadWriteCloser ga
	var rwc ReadWriteCloser = file
	rwc.Write("New Content")
	fmt.Printf("  ReadWriteCloser.Read(): %s\n", rwc.Read())
}

// demonstrateInterfaceNilCheck - Interface nil tekshirish
func demonstrateInterfaceNilCheck() {
	fmt.Println("Interface nil tekshirish:")

	var r Reader
	fmt.Printf("  r == nil: %t\n", r == nil)

	// Nil pointer interface ga uzatilsa
	var f *File
	r = f
	fmt.Printf("  r == nil (nil pointer): %t\n", r == nil) // False!

	// Nil tekshirish - type assertion
	if r != nil {
		if file, ok := r.(*File); ok && file != nil {
			fmt.Printf("  File is not nil\n")
		} else {
			fmt.Printf("  File is nil\n")
		}
	}

	// To'g'ri nil tekshirish
	var r2 Reader = (*File)(nil)
	if r2 == nil {
		fmt.Printf("  r2 is nil\n")
	} else {
		fmt.Printf("  r2 is not nil (nil pointer in interface)\n")
	}
}

// demonstrateInterfaceCompositionWithMethods - Interface composition bilan metodlar
func demonstrateInterfaceCompositionWithMethods() {
	fmt.Println("Interface composition bilan metodlar:")

	// ReadWriter - Reader va Writer metodlarini o'z ichiga oladi
	var rw ReadWriter = &File{content: "Initial"}

	// Reader metodlari
	fmt.Printf("  Reader.Read(): %s\n", rw.Read())

	// Writer metodlari
	rw.Write("Updated")
	fmt.Printf("  Writer.Write() dan keyin Read(): %s\n", rw.Read())

	// ReadWriteCloser - Reader, Writer, Closer metodlarini o'z ichiga oladi
	var rwc ReadWriteCloser = &File{content: "Content"}
	fmt.Printf("  ReadWriteCloser.Read(): %s\n", rwc.Read())
	rwc.Write("New Content")
	fmt.Printf("  ReadWriteCloser.Read() after Write: %s\n", rwc.Read())
	rwc.Close()
	fmt.Printf("  ReadWriteCloser.Read() after Close: %s\n", rwc.Read())
}

// demonstrateBuilderPattern - Builder pattern
func demonstrateBuilderPattern() {
	fmt.Println("Builder pattern:")

	// Builder struct
	type QueryBuilder struct {
		table  string
		where  string
		order  string
		limit  int
		offset int
	}

	// Builder metodlari
	table := func(name string) *QueryBuilder {
		return &QueryBuilder{table: name}
	}

	where := func(qb *QueryBuilder, condition string) *QueryBuilder {
		qb.where = condition
		return qb
	}

	orderBy := func(qb *QueryBuilder, field string) *QueryBuilder {
		qb.order = field
		return qb
	}

	limit := func(qb *QueryBuilder, n int) *QueryBuilder {
		qb.limit = n
		return qb
	}

	build := func(qb *QueryBuilder) string {
		query := "SELECT * FROM " + qb.table
		if qb.where != "" {
			query += " WHERE " + qb.where
		}
		if qb.order != "" {
			query += " ORDER BY " + qb.order
		}
		if qb.limit > 0 {
			query += fmt.Sprintf(" LIMIT %d", qb.limit)
		}
		return query
	}

	// Builder ishlatish
	qb := table("users")
	where(qb, "age > 18")
	orderBy(qb, "name")
	limit(qb, 10)
	query := build(qb)
	fmt.Printf("  Query: %s\n", query)
}

// SortStrategy interface - Strategy pattern uchun
type SortStrategy interface {
	Sort([]int) []int
}

// BubbleSort strategy
type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// QuickSort strategy (oddiy versiya)
type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	// Oddiy sort (aslida quick sort emas, lekin misol uchun)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

// Sorter - Context
type Sorter struct {
	strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(arr []int) []int {
	return s.strategy.Sort(arr)
}

// demonstrateStrategyPattern - Strategy pattern
func demonstrateStrategyPattern() {
	fmt.Println("Strategy pattern:")

	// Ishlatish
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	sorter := &Sorter{}
	sorter.SetStrategy(&BubbleSort{})
	sorted1 := sorter.Sort(numbers)
	fmt.Printf("  BubbleSort: %v\n", sorted1)

	sorter.SetStrategy(&QuickSort{})
	sorted2 := sorter.Sort(numbers)
	fmt.Printf("  QuickSort: %v\n", sorted2)
}

// Observer interface - Observer pattern uchun
type Observer interface {
	Update(string)
}

// Subject - Observer pattern uchun
type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.Notify()
}

// ConcreteObserver - Observer pattern uchun
type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(state string) {
	fmt.Printf("  Observer %s received: %s\n", co.name, state)
}

// demonstrateObserverPattern - Observer pattern
func demonstrateObserverPattern() {
	fmt.Println("Observer pattern:")

	// Ishlatish
	subject := &Subject{}
	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.SetState("State 1")
	subject.SetState("State 2")
}
