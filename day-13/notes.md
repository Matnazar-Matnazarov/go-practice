# Kun 13: Reflection (Refleksiya)

## Kirish

Go dasturlash tilida **Reflection** - bu runtime da type va value haqida ma'lumot olish imkoniyati. Reflection orqali type strukturasini, metodlarini va fieldlarini tekshirish mumkin.

---

## 1. Reflection Asoslari

### Reflection nima?

**Reflection** - bu runtime da type va value haqida ma'lumot olish mexanizmi.

**Reflection** - bu type va value ni tekshirish va manipulyatsiya qilish.

### Reflection nima uchun kerak?

1. **Type Inspection**
   - Runtime da type haqida ma'lumot olish
   - Type strukturasini tekshirish

2. **Dynamic Operations**
   - Type ni bilmasdan operatsiyalar bajarish
   - Generic kod yozish

3. **Serialization/Deserialization**
   - JSON, XML encoding/decoding
   - Database ORM

### Reflection vazifasi

Reflection quyidagi vazifalarni bajaradi:

1. **Type Inspection** - Type haqida ma'lumot olish
2. **Value Manipulation** - Value ni o'zgartirish
3. **Method Invocation** - Metodlarni chaqirish
4. **Field Access** - Fieldlarga kirish

### Reflection afzalliklari

✅ **Type Inspection** - Runtime da type haqida ma'lumot
✅ **Dynamic Operations** - Type ni bilmasdan operatsiyalar
✅ **Flexibility** - Kod qayta ishlatilishi

### Reflection kamchiliklari

❌ **Performance** - Runtime overhead
❌ **Type Safety** - Compile time da tekshirilmaydi
❌ **Murakkab** - Tushunish qiyin
❌ **Xavfli** - Panic qilishi mumkin

---

## 2. reflect Package

### reflect.Type va reflect.Value

**reflect.Type** - type haqida ma'lumot.

**reflect.Value** - value haqida ma'lumot.

```go
import "reflect"

// Type olish
var x int = 42
t := reflect.TypeOf(x)  // reflect.Type

// Value olish
v := reflect.ValueOf(x)  // reflect.Value
```

**Qoida:** `reflect.TypeOf()` va `reflect.ValueOf()` asosiy funksiyalar.

### TypeOf va ValueOf

**TypeOf** - type ni olish.

**ValueOf** - value ni olish.

```go
// TypeOf
t := reflect.TypeOf(42)        // int
t := reflect.TypeOf("hello")    // string

// ValueOf
v := reflect.ValueOf(42)        // reflect.Value
v := reflect.ValueOf("hello")  // reflect.Value
```

**Qoida:** TypeOf va ValueOf har qanday type uchun ishlaydi.

### Kind

**Kind** - type ning asosiy turi.

```go
v := reflect.ValueOf(42)
kind := v.Kind()  // reflect.Int

// Kind turlari
reflect.Int
reflect.String
reflect.Slice
reflect.Map
reflect.Struct
reflect.Ptr
reflect.Interface
```

**Kind** - type ning asosiy kategoriyasi.

---

## 3. Type Inspection

### Type Name va Kind

**Type Name** - type nomi.

**Kind** - type kategoriyasi.

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Ali", Age: 25}
t := reflect.TypeOf(p)

fmt.Println(t.Name())  // "Person"
fmt.Println(t.Kind())  // reflect.Struct
```

**Qoida:** Name - user-defined type nomi, Kind - asosiy kategoriya.

### Struct Fieldlar

**Struct fieldlar** - struct fieldlarini tekshirish.

```go
type Person struct {
    Name string
    Age  int
}

t := reflect.TypeOf(Person{})
for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Printf("Field: %s, Type: %s\n", field.Name, field.Type)
}
```

**NumField()** - fieldlar soni.

**Field(i)** - i-chi field.

### Field Tags

**Field tags** - struct field taglarini olish.

```go
type Person struct {
    Name string `json:"name" db:"person_name"`
    Age  int    `json:"age" db:"person_age"`
}

t := reflect.TypeOf(Person{})
field := t.Field(0)
tag := field.Tag.Get("json")  // "name"
```

**Tag** - field taglarini olish.

**Get()** - tag qiymatini olish.

### Methodlar

**Methodlar** - type metodlarini tekshirish.

```go
type Person struct {
    Name string
}

func (p Person) GetName() string {
    return p.Name
}

t := reflect.TypeOf(Person{})
for i := 0; i < t.NumMethod(); i++ {
    method := t.Method(i)
    fmt.Printf("Method: %s\n", method.Name)
}
```

**NumMethod()** - metodlar soni.

**Method(i)** - i-chi metod.

---

## 4. Value Manipulation

### Value Olish va O'zgartirish

**Value olish** - reflect.Value dan qiymat olish.

**Value o'zgartirish** - reflect.Value ni o'zgartirish.

```go
x := 42
v := reflect.ValueOf(&x)  // Pointer kerak
v = v.Elem()               // Pointer dan value ga
v.SetInt(100)              // Value ni o'zgartirish
```

**Qoida:** O'zgartirish uchun pointer kerak.

### Elem() - Pointer dan Value ga

**Elem()** - pointer dan value ga o'tish.

```go
x := 42
v := reflect.ValueOf(&x)  // Pointer
v = v.Elem()               // Value
v.SetInt(100)
```

**Qoida:** Pointer dan value ga o'tish uchun Elem() ishlatiladi.

### Set Methods

**Set methods** - value ni o'zgartirish metodlari.

```go
// Int
v.SetInt(100)

// String
v.SetString("hello")

// Bool
v.SetBool(true)

// Float
v.SetFloat(3.14)
```

**Qoida:** Har bir type uchun alohida Set metodi.

### CanSet() - O'zgartirish Mumkinligi

**CanSet()** - value ni o'zgartirish mumkinligini tekshirish.

```go
x := 42
v := reflect.ValueOf(x)     // Value (o'zgartirib bo'lmaydi)
fmt.Println(v.CanSet())   // false

v2 := reflect.ValueOf(&x)   // Pointer
v2 = v2.Elem()             // Value
fmt.Println(v2.CanSet())   // true
```

**Qoida:** CanSet() - o'zgartirish mumkinligini tekshiradi.

---

## 5. Struct Fieldlar bilan Ishlash

### Field Qiymatlarini Olish

**Field qiymatlari** - struct field qiymatlarini olish.

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Ali", Age: 25}
v := reflect.ValueOf(p)

name := v.Field(0).String()  // "Ali"
age := v.Field(1).Int()     // 25
```

**Field(i)** - i-chi field value.

**String()**, **Int()** - value ni olish.

### Field Qiymatlarini O'zgartirish

**Field o'zgartirish** - struct field qiymatlarini o'zgartirish.

```go
p := Person{Name: "Ali", Age: 25}
v := reflect.ValueOf(&p).Elem()

v.Field(0).SetString("Bob")
v.Field(1).SetInt(30)
```

**Qoida:** O'zgartirish uchun pointer va Elem() kerak.

### FieldByName

**FieldByName** - field nomi bo'yicha topish.

```go
p := Person{Name: "Ali", Age: 25}
v := reflect.ValueOf(&p).Elem()

nameField := v.FieldByName("Name")
nameField.SetString("Bob")
```

**FieldByName** - field nomi bo'yicha topish.

### FieldByIndex

**FieldByIndex** - nested struct fieldlar.

```go
type Address struct {
    City string
}

type Person struct {
    Name    string
    Address Address
}

p := Person{Name: "Ali", Address: Address{City: "Tashkent"}}
v := reflect.ValueOf(&p).Elem()

cityField := v.FieldByIndex([]int{1, 0})  // Address.City
cityField.SetString("Samarkand")
```

**FieldByIndex** - nested fieldlar uchun.

---

## 6. Methodlar bilan Ishlash

### Method Chaqirish

**Method chaqirish** - reflect orqali metod chaqirish.

```go
type Person struct {
    Name string
}

func (p Person) GetName() string {
    return p.Name
}

p := Person{Name: "Ali"}
v := reflect.ValueOf(p)

method := v.MethodByName("GetName")
result := method.Call(nil)
fmt.Println(result[0].String())  // "Ali"
```

**MethodByName** - metod nomi bo'yicha topish.

**Call()** - metodni chaqirish.

### Method Parametrlari

**Method parametrlari** - metodga parametr uzatish.

```go
func (p Person) Greet(name string) string {
    return fmt.Sprintf("Hello, %s! I'm %s", name, p.Name)
}

p := Person{Name: "Ali"}
v := reflect.ValueOf(p)

method := v.MethodByName("Greet")
args := []reflect.Value{reflect.ValueOf("Bob")}
result := method.Call(args)
fmt.Println(result[0].String())  // "Hello, Bob! I'm Ali"
```

**Call()** - parametrlar bilan metod chaqirish.

### Method Return Values

**Method return values** - metod qaytargan qiymatlar.

```go
method := v.MethodByName("GetName")
result := method.Call(nil)

// result - []reflect.Value
if len(result) > 0 {
    value := result[0].String()
    fmt.Println(value)
}
```

**Qoida:** Call() - []reflect.Value qaytaradi.

---

## 7. Slice va Map bilan Ishlash

### Slice bilan Ishlash

**Slice** - slice bilan reflection.

```go
slice := []int{1, 2, 3}
v := reflect.ValueOf(slice)

fmt.Println(v.Len())        // 3
fmt.Println(v.Cap())        // 3
fmt.Println(v.Index(0).Int())  // 1
```

**Len()** - slice uzunligi.

**Cap()** - slice sig'imi.

**Index(i)** - i-chi element.

### Map bilan Ishlash

**Map** - map bilan reflection.

```go
m := map[string]int{"a": 1, "b": 2}
v := reflect.ValueOf(m)

keys := v.MapKeys()
for _, key := range keys {
    value := v.MapIndex(key)
    fmt.Printf("%s: %d\n", key.String(), value.Int())
}
```

**MapKeys()** - map kalitlari.

**MapIndex()** - map qiymati.

### Slice va Map O'zgartirish

**Slice va Map o'zgartirish** - slice va map ni o'zgartirish.

```go
// Slice
slice := []int{1, 2, 3}
v := reflect.ValueOf(&slice).Elem()
v.Index(0).SetInt(10)

// Map
m := map[string]int{"a": 1}
v := reflect.ValueOf(&m).Elem()
key := reflect.ValueOf("b")
value := reflect.ValueOf(2)
v.SetMapIndex(key, value)
```

**SetMapIndex()** - map ga element qo'shish/o'zgartirish.

---

## 8. Type Assertion va Type Switch

### Type Assertion

**Type assertion** - reflect.Value dan type ga.

```go
v := reflect.ValueOf(42)

// Type assertion
if v.Kind() == reflect.Int {
    value := v.Int()
    fmt.Println(value)
}
```

**Qoida:** Kind() tekshirib, keyin value olish.

### Type Switch

**Type switch** - Kind bo'yicha switch.

```go
func processValue(v reflect.Value) {
    switch v.Kind() {
    case reflect.Int:
        fmt.Printf("Int: %d\n", v.Int())
    case reflect.String:
        fmt.Printf("String: %s\n", v.String())
    case reflect.Slice:
        fmt.Printf("Slice length: %d\n", v.Len())
    default:
        fmt.Printf("Unknown: %v\n", v.Interface())
    }
}
```

**Kind()** - type kategoriyasi.

**Interface()** - value ni interface{} ga o'tkazish.

---

## 9. Muhim Qoidalar

### Performance

**Performance** - reflection sekin.

```go
// Yomon: Har safar reflection
for i := 0; i < 1000; i++ {
    v := reflect.ValueOf(x)
    // ...
}

// Yaxshi: Bir marta reflection
v := reflect.ValueOf(x)
for i := 0; i < 1000; i++ {
    // ...
}
```

**Qoida:** Reflection ni kam ishlatish.

### Type Safety

**Type Safety** - compile time da tekshirilmaydi.

```go
// Xavfli: Panic qilishi mumkin
v := reflect.ValueOf(x)
v.SetInt(100)  // Agar x int bo'lmasa, panic!

// Xavfsiz: Tekshirish
if v.Kind() == reflect.Int && v.CanSet() {
    v.SetInt(100)
}
```

**Qoida:** Har doim tekshirish.

### Pointer va Value

**Pointer va Value** - o'zgartirish uchun pointer kerak.

```go
// Xato: O'zgartirib bo'lmaydi
x := 42
v := reflect.ValueOf(x)
v.SetInt(100)  // Panic!

// To'g'ri: Pointer orqali
x := 42
v := reflect.ValueOf(&x).Elem()
v.SetInt(100)  // OK
```

**Qoida:** O'zgartirish uchun pointer va Elem().

---

## Xulosa

### Reflection

**Nima:** Runtime da type va value haqida ma'lumot olish

**Vazifasi:**
- Type inspection
- Value manipulation
- Method invocation
- Field access

**Afzalliklari:**
- Type inspection
- Dynamic operations
- Flexibility

**Kamchiliklari:**
- Performance overhead
- Type safety yo'q
- Murakkab
- Xavfli

### Keyingi qadamlar

- Advanced reflection patterns
- JSON encoding/decoding
- Database ORM
- Configuration parsing
- Real-world examples
