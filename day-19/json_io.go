package main

import (
	"encoding/json"
	"fmt"
	"os"
)


// 2. JSON ENCODING/DECODING


// Person - JSON misoli uchun struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city,omitempty"`
}

// Employee - JSON tags misoli
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`              // JSON ga kiritmaslik
	Salary   int    `json:"salary,omitempty"` // Bo'sh bo'lsa, o'tkazib yuborish
	Active   bool   `json:"active"`
}

// demonstrateJSONMarshalUnmarshal - json.Marshal va json.Unmarshal
func demonstrateJSONMarshalUnmarshal() {
	// Struct yaratish
	person := Person{
		Name: "Ali",
		Age:  25,
		City: "Tashkent",
	}

	// JSON ga o'tkazish (Marshal)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("  Marshal xatosi: %v\n", err)
		return
	}
	fmt.Printf("  JSON (Marshal): %s\n", string(jsonData))

	// JSON dan o'tkazish (Unmarshal)
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Printf("  Unmarshal xatosi: %v\n", err)
		return
	}
	fmt.Printf("  Decoded Person: %+v\n", decodedPerson)
}

// demonstrateJSONMarshalIndent - json.MarshalIndent (formatlangan JSON)
func demonstrateJSONMarshalIndent() {
	person := Person{
		Name: "Bob",
		Age:  30,
		City: "Samarkand",
	}

	// Formatlangan JSON
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("  MarshalIndent xatosi: %v\n", err)
		return
	}
	fmt.Printf("  Formatlangan JSON:\n%s\n", string(jsonData))
}

// demonstrateJSONEncoderDecoder - json.Encoder va json.Decoder (stream)
func demonstrateJSONEncoderDecoder() {
	// Fayl yaratish
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Printf("  Fayl yaratish xatosi: %v\n", err)
		return
	}
	defer file.Close()

	// Encoder yaratish va yozish
	encoder := json.NewEncoder(file)
	person := Person{Name: "Charlie", Age: 35, City: "Bukhara"}
	err = encoder.Encode(person)
	if err != nil {
		fmt.Printf("  Encode xatosi: %v\n", err)
		return
	}
	fmt.Println("  person.json fayliga yozildi")

	// Faylni qayta ochish o'qish uchun
	file.Close()
	file, err = os.Open("person.json")
	if err != nil {
		fmt.Printf("  Fayl ochish xatosi: %v\n", err)
		return
	}
	defer file.Close()

	// Decoder yaratish va o'qish
	decoder := json.NewDecoder(file)
	var decodedPerson Person
	err = decoder.Decode(&decodedPerson)
	if err != nil {
		fmt.Printf("  Decode xatosi: %v\n", err)
		return
	}
	fmt.Printf("  O'qilgan Person: %+v\n", decodedPerson)
}

// demonstrateJSONTags - JSON tags
func demonstrateJSONTags() {
	// Employee yaratish
	employee := Employee{
		ID:       1,
		Name:     "David",
		Email:    "david@example.com",
		Password: "secret123", // Bu JSON ga kiritilmaydi
		Salary:   5000,
		Active:   true,
	}

	// JSON ga o'tkazish
	jsonData, err := json.MarshalIndent(employee, "", "  ")
	if err != nil {
		fmt.Printf("  Marshal xatosi: %v\n", err)
		return
	}
	fmt.Printf("  Employee JSON (Password ko'rinmaydi):\n%s\n", string(jsonData))

	// omitempty misoli
	employee2 := Employee{
		ID:       2,
		Name:     "Eve",
		Email:    "eve@example.com",
		Password: "secret456",
		// Salary o'tkazib yuborilgan (0 qiymat)
		Active: true,
	}

	jsonData2, err := json.MarshalIndent(employee2, "", "  ")
	if err != nil {
		fmt.Printf("  Marshal xatosi: %v\n", err)
		return
	}
	fmt.Printf("  Employee JSON (Salary omitempty):\n%s\n", string(jsonData2))
}


// QO'SHIMCHA: JSON Fayl Operatsiyalari


// SavePersonToFile - Person ni JSON faylga saqlash
func SavePersonToFile(filename string, person Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(person)
}

// LoadPersonFromFile - Person ni JSON fayldan yuklash
func LoadPersonFromFile(filename string) (*Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var person Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}
