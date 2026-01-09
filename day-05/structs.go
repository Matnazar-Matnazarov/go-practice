package main

import "fmt"

// Person - oddiy struct
type Person struct {
	Name string
	Age  int
	City string
}

// Student - struct metodlar bilan
type Student struct {
	Person
	StudentID int
	Grade     string
}

// Employee - struct metodlar bilan
type Employee struct {
	Person
	EmployeeID int
	Salary     float64
	Department string
}

// Rectangle - geometrik shakl
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle - geometrik shakl
type Circle struct {
	Radius float64
}

// Book - kitob struktura
type Book struct {
	Title  string
	Author string
	Pages  int
	Price  float64
}

// Address - manzil struktura
type Address struct {
	Street   string
	City     string
	Country  string
	PostCode string
}

// Contact - aloqa ma'lumotlari
type Contact struct {
	Email string
	Phone string
}

// User - foydalanuvchi (nested struct)
type User struct {
	Person
	Address Address
	Contact Contact
}

// demonstrateStructBasic - struct asoslari
func demonstrateStructBasic() {
	// Struct literal sintaksis
	person1 := Person{
		Name: "Matnazar",
		Age:  25,
		City: "Toshkent",
	}

	fmt.Println("Person 1:", person1)
	fmt.Printf("Ism: %s, Yosh: %d, Shahar: %s\n", person1.Name, person1.Age, person1.City)

	// Qisqa sintaksis
	person2 := Person{"Ali", 30, "Samarqand"}
	fmt.Println("Person 2:", person2)

	// Field nomlarsiz
	person3 := Person{
		"Vali",
		28,
		"Buxoro",
	}
	fmt.Println("Person 3:", person3)
}

// demonstrateStructAccess - struct fieldlarga kirish
func demonstrateStructAccess() {
	person := Person{
		Name: "Matnazar",
		Age:  25,
		City: "Toshkent",
	}

	// Fieldlarga kirish
	fmt.Printf("Ism: %s\n", person.Name)
	fmt.Printf("Yosh: %d\n", person.Age)
	fmt.Printf("Shahar: %s\n", person.City)

	// Fieldlarni o'zgartirish
	person.Age = 26
	person.City = "Samarqand"
	fmt.Println("O'zgartirilgan:", person)
}

// demonstrateStructPointer - struct pointer
func demonstrateStructPointer() {
	person := &Person{
		Name: "Matnazar",
		Age:  25,
		City: "Toshkent",
	}

	fmt.Println("Pointer orqali:", person)
	fmt.Printf("Ism: %s\n", person.Name) // Avtomatik dereference

	// Pointer orqali o'zgartirish
	person.Age = 26
	fmt.Println("O'zgartirilgan:", person)
}

// demonstrateStructEmbedding - struct embedding (inheritance)
func demonstrateStructEmbedding() {
	student := Student{
		Person: Person{
			Name: "Matnazar",
			Age:  25,
			City: "Toshkent",
		},
		StudentID: 12345,
		Grade:     "A",
	}

	fmt.Println("Student:", student)
	fmt.Printf("Ism: %s, StudentID: %d, Grade: %s\n", student.Name, student.StudentID, student.Grade)
}

// demonstrateNestedStruct - nested struct
func demonstrateNestedStruct() {
	user := User{
		Person: Person{
			Name: "Matnazar",
			Age:  25,
			City: "Toshkent",
		},
		Address: Address{
			Street:   "Amir Temur ko'chasi",
			City:     "Toshkent",
			Country:  "O'zbekiston",
			PostCode: "100000",
		},
		Contact: Contact{
			Email: "matnazar@example.com",
			Phone: "+998901234567",
		},
	}

	fmt.Println("User ma'lumotlari:")
	fmt.Printf("Ism: %s\n", user.Name)
	fmt.Printf("Manzil: %s, %s\n", user.Address.Street, user.Address.City)
	fmt.Printf("Email: %s\n", user.Contact.Email)
}

// demonstrateStructArray - struct array
func demonstrateStructArray() {
	people := [3]Person{
		{"Matnazar", 25, "Toshkent"},
		{"Ali", 30, "Samarqand"},
		{"Vali", 28, "Buxoro"},
	}

	fmt.Println("People array:")
	for i, person := range people {
		fmt.Printf("  %d. %s, %d yosh, %s\n", i+1, person.Name, person.Age, person.City)
	}
}

// demonstrateStructSlice - struct slice
func demonstrateStructSlice() {
	students := []Student{
		{
			Person:    Person{"Matnazar", 25, "Toshkent"},
			StudentID: 1,
			Grade:     "A",
		},
		{
			Person:    Person{"Ali", 30, "Samarqand"},
			StudentID: 2,
			Grade:     "B",
		},
		{
			Person:    Person{"Vali", 28, "Buxoro"},
			StudentID: 3,
			Grade:     "A",
		},
	}

	fmt.Println("Students slice:")
	for _, student := range students {
		fmt.Printf("  %s (ID: %d, Grade: %s)\n", student.Name, student.StudentID, student.Grade)
	}
}

// demonstrateStructMap - struct map
func demonstrateStructMap() {
	employees := map[int]Employee{
		1: {
			Person:     Person{"Matnazar", 25, "Toshkent"},
			EmployeeID: 1,
			Salary:     5000.0,
			Department: "IT",
		},
		2: {
			Person:     Person{"Ali", 30, "Samarqand"},
			EmployeeID: 2,
			Salary:     6000.0,
			Department: "HR",
		},
	}

	fmt.Println("Employees map:")
	for id, emp := range employees {
		fmt.Printf("  ID %d: %s, %s, %.2f\n", id, emp.Name, emp.Department, emp.Salary)
	}
}

// Rectangle metodlari

// Area - to'rtburchak yuzasi
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter - to'rtburchak perimetri
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle metodlari

// Area - doira yuzasi
func (c Circle) Area() float64 {
	const pi = 3.14159
	return pi * c.Radius * c.Radius
}

// Perimeter - doira perimetri
func (c Circle) Perimeter() float64 {
	const pi = 3.14159
	return 2 * pi * c.Radius
}

// Person metodlari

// GetInfo - person ma'lumotlarini olish
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s, %d yosh, %s", p.Name, p.Age, p.City)
}

// IsAdult - katta yoshni tekshirish
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Student metodlari

// GetStudentInfo - student ma'lumotlarini olish
func (s Student) GetStudentInfo() string {
	return fmt.Sprintf("%s (ID: %d, Grade: %s)", s.Name, s.StudentID, s.Grade)
}

// Employee metodlari

// GetEmployeeInfo - employee ma'lumotlarini olish
func (e Employee) GetEmployeeInfo() string {
	return fmt.Sprintf("%s (ID: %d, %s, %.2f)", e.Name, e.EmployeeID, e.Department, e.Salary)
}

// CalculateAnnualSalary - yillik maosh
func (e Employee) CalculateAnnualSalary() float64 {
	return e.Salary * 12
}

// Book metodlari

// GetBookInfo - kitob ma'lumotlari
func (b Book) GetBookInfo() string {
	return fmt.Sprintf("'%s' - %s, %d sahifa, %.2f so'm", b.Title, b.Author, b.Pages, b.Price)
}

// IsExpensive - qimmat kitob
func (b Book) IsExpensive() bool {
	return b.Price > 50000
}

// demonstrateStructMethods - struct metodlar
func demonstrateStructMethods() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("To'rtburchak: %.2f x %.2f\n", rect.Width, rect.Height)
	fmt.Printf("Yuzasi: %.2f\n", rect.Area())
	fmt.Printf("Perimetri: %.2f\n", rect.Perimeter())

	circle := Circle{Radius: 5}
	fmt.Printf("\nDoira: radius=%.2f\n", circle.Radius)
	fmt.Printf("Yuzasi: %.2f\n", circle.Area())
	fmt.Printf("Perimetri: %.2f\n", circle.Perimeter())
}

// demonstratePointerReceiver - pointer receiver metodlar
func demonstratePointerReceiver() {
	person := Person{
		Name: "Matnazar",
		Age:  25,
		City: "Toshkent",
	}

	fmt.Println("Boshlang'ich:", person.GetInfo())

	// Value receiver - o'zgarmaydi
	person.UpdateAge(26)
	fmt.Println("Value receiver dan keyin:", person.GetInfo())

	// Pointer receiver - o'zgaradi
	person.UpdateAgePointer(27)
	fmt.Println("Pointer receiver dan keyin:", person.GetInfo())
}

// UpdateAge - value receiver (o'zgarmaydi, misol uchun)
// Note: Value receiver struct nusxasida ishlaydi, asl struct o'zgarmaydi
func (p Person) UpdateAge(newAge int) {
	p.Age = newAge // nolint:staticcheck // misol uchun value receiver
}

// UpdateAgePointer - pointer receiver (o'zgaradi)
func (p *Person) UpdateAgePointer(newAge int) {
	p.Age = newAge
}

// demonstrateStructComparison - struct solishtirish
func demonstrateStructComparison() {
	person1 := Person{"Matnazar", 25, "Toshkent"}
	person2 := Person{"Matnazar", 25, "Toshkent"}
	person3 := Person{"Ali", 30, "Samarqand"}

	fmt.Printf("person1 == person2: %t\n", person1 == person2)
	fmt.Printf("person1 == person3: %t\n", person1 == person3)
}

// demonstrateStructAsParameter - struct parametr sifatida
func demonstrateStructAsParameter() {
	person := Person{"Matnazar", 25, "Toshkent"}
	printPerson(person)
	printPersonPointer(&person)
}

// printPerson - struct value parametr
func printPerson(p Person) {
	fmt.Printf("Person: %s, %d yosh, %s\n", p.Name, p.Age, p.City)
}

// printPersonPointer - struct pointer parametr
func printPersonPointer(p *Person) {
	fmt.Printf("Person (pointer): %s, %d yosh, %s\n", p.Name, p.Age, p.City)
}

// demonstrateStructReturn - struct qaytarish
func demonstrateStructReturn() {
	person := createPerson("Matnazar", 25, "Toshkent")
	fmt.Println("Yaratilgan person:", person.GetInfo())
}

// createPerson - struct yaratish va qaytarish
func createPerson(name string, age int, city string) Person {
	return Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// demonstrateStructFilter - struct slice filtrlash
func demonstrateStructFilter() {
	people := []Person{
		{"Matnazar", 25, "Toshkent"},
		{"Ali", 17, "Samarqand"},
		{"Vali", 30, "Buxoro"},
		{"Bobur", 16, "Andijon"},
	}

	adults := filterAdults(people)
	fmt.Println("Kattalar:")
	for _, person := range adults {
		fmt.Printf("  %s, %d yosh\n", person.Name, person.Age)
	}
}

// filterAdults - kattalarni filtrlash
func filterAdults(people []Person) []Person {
	var result []Person
	for _, person := range people {
		if person.IsAdult() {
			result = append(result, person)
		}
	}
	return result
}

// demonstrateStructSort - struct slice sort qilish
func demonstrateStructSort() {
	students := []Student{
		{Person: Person{"Matnazar", 25, "Toshkent"}, StudentID: 3, Grade: "A"},
		{Person: Person{"Ali", 30, "Samarqand"}, StudentID: 1, Grade: "B"},
		{Person: Person{"Vali", 28, "Buxoro"}, StudentID: 2, Grade: "A"},
	}

	fmt.Println("Talabalar (ID bo'yicha):")
	for _, student := range students {
		fmt.Printf("  ID: %d, %s\n", student.StudentID, student.Name)
	}
}
