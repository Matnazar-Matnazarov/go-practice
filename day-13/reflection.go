package main

import (
	"fmt"
	"reflect"
)


// 1. REFLECTION ASOSLARI


// demonstrateTypeOfAndValueOf - TypeOf va ValueOf
func demonstrateTypeOfAndValueOf() {
	// TypeOf
	var x int = 42
	t := reflect.TypeOf(x)
	fmt.Printf("  TypeOf(42): %s\n", t)

	var s string = "hello"
	t2 := reflect.TypeOf(s)
	fmt.Printf("  TypeOf(\"hello\"): %s\n", t2)

	// ValueOf
	v := reflect.ValueOf(x)
	fmt.Printf("  ValueOf(42): %v (kind: %s)\n", v, v.Kind())

	v2 := reflect.ValueOf(s)
	fmt.Printf("  ValueOf(\"hello\"): %v (kind: %s)\n", v2, v2.Kind())
}

// demonstrateKind - Kind
func demonstrateKind() {
	values := []interface{}{
		42,
		"hello",
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		true,
		3.14,
	}

	for _, val := range values {
		v := reflect.ValueOf(val)
		fmt.Printf("  %v: kind = %s\n", val, v.Kind())
	}
}


// 2. TYPE INSPECTION


// Person - misol struct
type Person struct {
	Name string
	Age  int
}

// GetName - Person metod
func (p Person) GetName() string {
	return p.Name
}

// GetAge - Person metod
func (p Person) GetAge() int {
	return p.Age
}

// demonstrateTypeNameAndKind - Type name va kind
func demonstrateTypeNameAndKind() {
	p := Person{Name: "Ali", Age: 25}
	t := reflect.TypeOf(p)

	fmt.Printf("  Type name: %s\n", t.Name())
	fmt.Printf("  Type kind: %s\n", t.Kind())
}

// demonstrateStructFields - Struct fieldlar
func demonstrateStructFields() {
	t := reflect.TypeOf(Person{})
	fmt.Printf("  Struct: %s\n", t.Name())
	fmt.Printf("  Fieldlar soni: %d\n", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("    Field %d: %s (type: %s)\n", i, field.Name, field.Type)
	}
}

// PersonWithTags - taglar bilan struct
type PersonWithTags struct {
	Name string `json:"name" db:"person_name"`
	Age  int    `json:"age" db:"person_age"`
}

// demonstrateFieldTags - Field tags
func demonstrateFieldTags() {
	t := reflect.TypeOf(PersonWithTags{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		fmt.Printf("  Field: %s\n", field.Name)
		fmt.Printf("    json tag: %s\n", jsonTag)
		fmt.Printf("    db tag: %s\n", dbTag)
	}
}

// demonstrateMethods - Methodlar
func demonstrateMethods() {
	t := reflect.TypeOf(Person{})
	fmt.Printf("  Methodlar soni: %d\n", t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("    Method %d: %s\n", i, method.Name)
	}
}


// 3. VALUE MANIPULATION


// demonstrateValueGetAndSet - Value olish va o'zgartirish
func demonstrateValueGetAndSet() {
	// Value olish
	x := 42
	v := reflect.ValueOf(x)
	fmt.Printf("  Original value: %d\n", v.Int())

	// Value o'zgartirish (pointer kerak)
	x2 := 42
	v2 := reflect.ValueOf(&x2).Elem() // Pointer dan value ga
	v2.SetInt(100)
	fmt.Printf("  After SetInt(100): %d\n", x2)
}

// demonstrateCanSet - CanSet
func demonstrateCanSet() {
	x := 42

	// Value (o'zgartirib bo'lmaydi)
	v := reflect.ValueOf(x)
	fmt.Printf("  ValueOf(x).CanSet(): %t\n", v.CanSet())

	// Pointer (o'zgartirish mumkin)
	v2 := reflect.ValueOf(&x).Elem()
	fmt.Printf("  ValueOf(&x).Elem().CanSet(): %t\n", v2.CanSet())

	// O'zgartirish
	if v2.CanSet() {
		v2.SetInt(100)
		fmt.Printf("  After SetInt(100): x = %d\n", x)
	}
}


// 4. STRUCT FIELDLAR BILAN ISHLASH


// demonstrateGetFieldValues - Field qiymatlarini olish
func demonstrateGetFieldValues() {
	p := Person{Name: "Ali", Age: 25}
	v := reflect.ValueOf(p)

	fmt.Printf("  Person: %+v\n", p)
	fmt.Printf("  Field(0) (Name): %s\n", v.Field(0).String())
	fmt.Printf("  Field(1) (Age): %d\n", int(v.Field(1).Int()))
}

// demonstrateSetFieldValues - Field qiymatlarini o'zgartirish
func demonstrateSetFieldValues() {
	p := Person{Name: "Ali", Age: 25}
	fmt.Printf("  Before: %+v\n", p)

	v := reflect.ValueOf(&p).Elem()
	v.Field(0).SetString("Bob")
	v.Field(1).SetInt(30)

	fmt.Printf("  After: %+v\n", p)
}

// demonstrateFieldByName - FieldByName
func demonstrateFieldByName() {
	p := Person{Name: "Ali", Age: 25}
	v := reflect.ValueOf(&p).Elem()

	// FieldByName orqali
	nameField := v.FieldByName("Name")
	if nameField.IsValid() {
		fmt.Printf("  FieldByName(\"Name\"): %s\n", nameField.String())
		nameField.SetString("Bob")
		fmt.Printf("  After SetString(\"Bob\"): %+v\n", p)
	}

	ageField := v.FieldByName("Age")
	if ageField.IsValid() {
		fmt.Printf("  FieldByName(\"Age\"): %d\n", int(ageField.Int()))
	}
}


// 5. METHODLAR BILAN ISHLASH


// demonstrateMethodCall - Method chaqirish
func demonstrateMethodCall() {
	p := Person{Name: "Ali", Age: 25}
	v := reflect.ValueOf(p)

	// GetName metodini chaqirish
	method := v.MethodByName("GetName")
	if method.IsValid() {
		result := method.Call(nil)
		if len(result) > 0 {
			fmt.Printf("  GetName() result: %s\n", result[0].String())
		}
	}

	// GetAge metodini chaqirish
	method2 := v.MethodByName("GetAge")
	if method2.IsValid() {
		result := method2.Call(nil)
		if len(result) > 0 {
			fmt.Printf("  GetAge() result: %d\n", int(result[0].Int()))
		}
	}
}

// Greet - Person metod (parametr bilan)
func (p Person) Greet(name string) string {
	return fmt.Sprintf("Hello, %s! I'm %s", name, p.Name)
}

// demonstrateMethodWithParameters - Method parametrlari
func demonstrateMethodWithParameters() {
	p := Person{Name: "Ali", Age: 25}
	v := reflect.ValueOf(p)

	// Greet metodini chaqirish
	method := v.MethodByName("Greet")
	if method.IsValid() {
		args := []reflect.Value{reflect.ValueOf("Bob")}
		result := method.Call(args)
		if len(result) > 0 {
			fmt.Printf("  Greet(\"Bob\"): %s\n", result[0].String())
		}
	}
}


// 6. SLICE VA MAP BILAN ISHLASH


// demonstrateSliceReflection - Slice bilan ishlash
func demonstrateSliceReflection() {
	slice := []int{1, 2, 3, 4, 5}
	v := reflect.ValueOf(slice)

	fmt.Printf("  Slice: %v\n", slice)
	fmt.Printf("  Len(): %d\n", v.Len())
	fmt.Printf("  Cap(): %d\n", v.Cap())

	// Elementlarni olish
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("    Index(%d): %d\n", i, v.Index(i).Int())
	}
}

// demonstrateMapReflection - Map bilan ishlash
func demonstrateMapReflection() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 10,
	}
	v := reflect.ValueOf(m)

	fmt.Printf("  Map: %v\n", m)

	// Kalitlar
	keys := v.MapKeys()
	fmt.Printf("  Keys: ")
	for _, key := range keys {
		fmt.Printf("%s ", key.String())
	}
	fmt.Println()

	// Qiymatlar
	for _, key := range keys {
		value := v.MapIndex(key)
		fmt.Printf("    %s: %d\n", key.String(), int(value.Int()))
	}
}


// 7. TYPE ASSERTION VA TYPE SWITCH


// demonstrateTypeAssertion - Type assertion
func demonstrateTypeAssertion() {
	values := []interface{}{
		42,
		"hello",
		true,
		3.14,
	}

	for _, val := range values {
		v := reflect.ValueOf(val)
		fmt.Printf("  Value: %v\n", val)

		switch v.Kind() {
		case reflect.Int:
			fmt.Printf("    Type: int, Value: %d\n", v.Int())
		case reflect.String:
			fmt.Printf("    Type: string, Value: %s\n", v.String())
		case reflect.Bool:
			fmt.Printf("    Type: bool, Value: %t\n", v.Bool())
		case reflect.Float64:
			fmt.Printf("    Type: float64, Value: %.2f\n", v.Float())
		}
	}
}

// processValue - Type switch funksiyasi
func processValue(v reflect.Value) {
	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("    Int: %d\n", v.Int())
	case reflect.String:
		fmt.Printf("    String: %s\n", v.String())
	case reflect.Bool:
		fmt.Printf("    Bool: %t\n", v.Bool())
	case reflect.Float64:
		fmt.Printf("    Float64: %.2f\n", v.Float())
	case reflect.Slice:
		fmt.Printf("    Slice length: %d\n", v.Len())
	case reflect.Map:
		fmt.Printf("    Map keys: %d\n", len(v.MapKeys()))
	default:
		fmt.Printf("    Unknown: %v\n", v.Interface())
	}
}

// demonstrateTypeSwitch - Type switch
func demonstrateTypeSwitch() {
	values := []interface{}{
		42,
		"hello",
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		true,
		3.14,
	}

	for _, val := range values {
		v := reflect.ValueOf(val)
		fmt.Printf("  Processing: %v\n", val)
		processValue(v)
	}
}
