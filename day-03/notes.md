# Kun 3: Shartli operatorlar, Array va Slice

## Shartli operatorlar (Conditional Statements)

### if/else operatori

Shartli operatorlar dasturda qaror qabul qilish uchun ishlatiladi.

```go
if shart {
    // kod
} else {
    // kod
}
```

**Misol:**
```go
if age >= 18 {
    fmt.Println("Kattalar uchun")
} else {
    fmt.Println("Yosh bolalar uchun")
}
```

### if/else if/else zanjiri

Bir nechta shartlarni tekshirish:

```go
if shart1 {
    // kod
} else if shart2 {
    // kod
} else {
    // kod
}
```

### Shartli operatorlar xususiyatlari

1. **Murakkab shartlar** - `&&` (AND), `||` (OR), `!` (NOT)
2. **Qoldiq tekshirish** - `%` operatori
3. **Oraliq tekshirish** - `>=`, `<=`, `>`, `<`
4. **Tenglik tekshirish** - `==`, `!=`

### Switch operatori

Ko'p shartlarni tekshirish uchun qulay:

```go
switch o'zgaruvchi {
case qiymat1:
    // kod
case qiymat2:
    // kod
default:
    // kod
}
```

**Switch turlari:**
- **Oddiy switch** - o'zgaruvchi qiymatiga qarab
- **Switch expression** - shartga qarab
- **Switch type** - ma'lumot turiga qarab
- **Bir nechta case** - bir nechta qiymatlar

## Array (Massivlar)

### Array nima?

Array - bir xil tipdagi elementlarning ketma-ket joylashgan to'plami. 
Array uzunligi o'zgarmas (fixed size).

### Array e'lon qilish

```go
// Usul 1: var bilan
var numbers [5]int

// Usul 2: Literal sintaksis
numbers := [5]int{1, 2, 3, 4, 5}

// Usul 3: Avtomatik uzunlik
numbers := [...]int{1, 2, 3, 4, 5}
```

### Array operatsiyalari

- **Elementlarga kirish**: `arr[index]`
- **Elementni o'zgartirish**: `arr[index] = value`
- **Uzunlik**: `len(arr)`
- **Iteratsiya**: `for` loop yoki `range`

### Array xususiyatlari

1. **O'zgarmas uzunlik** - yaratilgandan keyin o'zgarmaydi
2. **Bir xil tip** - barcha elementlar bir xil tipda
3. **Index** - 0 dan boshlanadi
4. **Value type** - nusxalashda to'liq nusxa olinadi

## Slice (Kesmalar)

### Slice nima?

Slice - array ning dinamik ko'rinishi. 
Slice uzunligi o'zgarishi mumkin (dynamic size).

### Slice e'lon qilish

```go
// Usul 1: var bilan
var numbers []int

// Usul 2: Literal sintaksis
numbers := []int{1, 2, 3, 4, 5}

// Usul 3: make() bilan
numbers := make([]int, 3, 5) // uzunlik: 3, capacity: 5

// Usul 4: Array dan slice
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // [2, 3, 4]
```

### Slice operatsiyalari

- **append()** - element qo'shish
- **copy()** - slice nusxalash
- **len()** - uzunlik
- **cap()** - capacity (sig'im)
- **Slicing** - `slice[start:end]`

### Slice xususiyatlari

1. **Dinamik uzunlik** - append() bilan o'zgaradi
2. **Reference type** - nusxalashda reference olinadi
3. **Capacity** - ichki sig'im
4. **Flexible** - array dan yaxshiroq

### Array va Slice farqi

| Xususiyat | Array | Slice |
|-----------|-------|-------|
| Uzunlik | O'zgarmas | O'zgaruvchan |
| Tip | Value type | Reference type |
| E'lon | `[5]int` | `[]int` |
| make() | Ishlatilmaydi | Ishlatiladi |
| append() | Ishlatilmaydi | Ishlatiladi |

## Ko'p o'lchamli array va slice

### 2D Array

```go
var matrix [3][3]int
matrix[0] = [3]int{1, 2, 3}
matrix[1] = [3]int{4, 5, 6}
matrix[2] = [3]int{7, 8, 9}
```

### 2D Slice

```go
slice2D := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

## Amaliy misollar

### Array yig'indisi

```go
func sumArray(arr [5]int) int {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    return sum
}
```

### Slice filtrlash

```go
func filterEven(numbers []int) []int {
    var result []int
    for _, value := range numbers {
        if value%2 == 0 {
            result = append(result, value)
        }
    }
    return result
}
```

### Eng katta/kichik son

```go
func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    max := numbers[0]
    for _, value := range numbers {
        if value > max {
            max = value
        }
    }
    return max
}
```

## Bugun o'rganilganlar

✅ if/else shartli operatorlar
✅ if/else if/else zanjiri
✅ Switch operatori (oddiy, expression, type)
✅ Array (massivlar) - e'lon qilish, ishlatish
✅ Slice (kesmalar) - yaratish, append, copy
✅ Array va Slice operatsiyalari
✅ Ko'p o'lchamli array va slice
✅ Amaliy misollar va algoritmlar

## Keyingi kun

Kun 4: Map (xarita), Struct (struktura) va metodlar
