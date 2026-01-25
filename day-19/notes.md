# Kun 19: File I/O va JSON (Fayl Operatsiyalari va JSON)

## Kirish

Go dasturlash tilida **File I/O** va **JSON** - bu fayllar bilan ishlash va ma'lumotlarni serializatsiya/deserializatsiya qilishning asosiy mexanizmi. Bu real-world loyihalarda juda muhim.

---

## 1. File I/O Asoslari

### File I/O nima?

**File I/O** - bu fayllarni o'qish va yozish operatsiyalari.

**File I/O** - bu ma'lumotlarni faylda saqlash va o'qish.

### File I/O nima uchun kerak?

1. **Ma'lumotlarni saqlash**
   - Configuration fayllar
   - Log fayllar
   - Data fayllar

2. **Ma'lumotlarni almashish**
   - Import/Export
   - Backup/Restore
   - Data migration

### File I/O vazifasi

File I/O quyidagi vazifalarni bajaradi:

1. **Fayl yaratish** - Yangi fayl yaratish
2. **Fayl o'qish** - Fayl ma'lumotlarini o'qish
3. **Fayl yozish** - Faylga ma'lumot yozish
4. **Fayl o'chirish** - Faylni o'chirish

### File I/O afzalliklari

✅ **Ma'lumotlarni saqlash** - Persistent storage
✅ **Ma'lumotlarni almashish** - Data exchange
✅ **Configuration** - Sozlamalarni saqlash

---

## 2. os Package

### os.Open() - Fayl O'qish

**os.Open()** - faylni o'qish uchun ochish.

```go
file, err := os.Open("data.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

**Qoida:** `defer file.Close()` - faylni yopish.

### os.Create() - Fayl Yaratish

**os.Create()** - yangi fayl yaratish.

```go
file, err := os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

**Qoida:** `os.Create()` - fayl mavjud bo'lsa, o'chirib qayta yaratadi.

### os.OpenFile() - Fayl Ochish (Kengaytirilgan)

**os.OpenFile()** - faylni ochish (flag va mode bilan).

```go
file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

**Qoida:** `os.OpenFile()` - flag va mode bilan fayl ochish.

---

## 3. io Package

### io.ReadAll() - Barcha Ma'lumotni O'qish

**io.ReadAll()** - faylning barcha ma'lumotini o'qish.

```go
data, err := io.ReadAll(file)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `io.ReadAll()` - barcha ma'lumotni bir marta o'qish.

### io.Copy() - Ma'lumotni Nusxalash

**io.Copy()** - ma'lumotni bir joydan ikkinchisiga nusxalash.

```go
_, err := io.Copy(dst, src)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `io.Copy()` - ma'lumotni nusxalash.

### io.WriteString() - String Yozish

**io.WriteString()** - string ni faylga yozish.

```go
_, err := io.WriteString(file, "Hello, World!")
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `io.WriteString()` - string ni faylga yozish.

---

## 4. bufio Package

### bufio.Scanner - Qatorma-qator O'qish

**bufio.Scanner** - faylni qatorma-qator o'qish.

```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
```

**Qoida:** `bufio.Scanner` - katta fayllarni qatorma-qator o'qish.

### bufio.Writer - Buffered Yozish

**bufio.Writer** - buffered yozish.

```go
writer := bufio.NewWriter(file)
writer.WriteString("Hello\n")
writer.Flush()
```

**Qoida:** `bufio.Writer` - tezroq yozish uchun buffer.

---

## 5. ioutil Package (Deprecated)

### ioutil.ReadFile() - Fayl O'qish

**ioutil.ReadFile()** - faylni bir marta o'qish (deprecated).

```go
data, err := ioutil.ReadFile("data.txt")
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `os.ReadFile()` ishlatish kerak (Go 1.16+).

### ioutil.WriteFile() - Fayl Yozish

**ioutil.WriteFile()** - faylga yozish (deprecated).

```go
err := ioutil.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `os.WriteFile()` ishlatish kerak (Go 1.16+).

---

## 6. os.ReadFile() va os.WriteFile()

### os.ReadFile() - Fayl O'qish

**os.ReadFile()** - faylni bir marta o'qish.

```go
data, err := os.ReadFile("data.txt")
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `os.ReadFile()` - kichik fayllar uchun.

### os.WriteFile() - Fayl Yozish

**os.WriteFile()** - faylga yozish.

```go
data := []byte("Hello, World!")
err := os.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `os.WriteFile()` - kichik fayllar uchun.

---

## 7. JSON Encoding/Decoding

### JSON nima?

**JSON** - JavaScript Object Notation, ma'lumotlarni saqlash va almashish formati.

**JSON** - bu text-based format, odamlar o'qishi oson.

### json.Marshal() - JSON ga O'tkazish

**json.Marshal()** - Go struct ni JSON ga o'tkazish.

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

person := Person{Name: "Ali", Age: 25}
data, err := json.Marshal(person)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `json.Marshal()` - struct ni JSON byte array ga o'tkazadi.

### json.Unmarshal() - JSON dan O'tkazish

**json.Unmarshal()** - JSON ni Go struct ga o'tkazish.

```go
var person Person
err := json.Unmarshal(data, &person)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `json.Unmarshal()` - JSON ni struct ga o'tkazadi.

### json.MarshalIndent() - Formatlangan JSON

**json.MarshalIndent()** - formatlangan JSON.

```go
data, err := json.MarshalIndent(person, "", "  ")
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `json.MarshalIndent()` - chiroyli formatlangan JSON.

---

## 8. JSON Encoder/Decoder

### json.Encoder - Stream Encoding

**json.Encoder** - stream orqali JSON encoding.

```go
encoder := json.NewEncoder(file)
err := encoder.Encode(person)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `json.Encoder` - katta ma'lumotlar uchun.

### json.Decoder - Stream Decoding

**json.Decoder** - stream orqali JSON decoding.

```go
decoder := json.NewDecoder(file)
var person Person
err := decoder.Decode(&person)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `json.Decoder` - katta ma'lumotlar uchun.

---

## 9. JSON Tags

### JSON Field Tags

**JSON tags** - struct field taglar.

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
    City string `json:"-"`
}
```

**Qoida:**
- `json:"name"` - field nomi
- `json:"age,omitempty"` - bo'sh bo'lsa, o'tkazib yuborish
- `json:"-"` - JSON ga kiritmaslik

---

## 10. File Operations

### File Mavjudligini Tekshirish

**os.Stat()** - fayl haqida ma'lumot olish.

```go
info, err := os.Stat("data.txt")
if os.IsNotExist(err) {
    // Fayl mavjud emas
}
```

**Qoida:** `os.Stat()` - fayl haqida ma'lumot.

### File O'chirish

**os.Remove()** - faylni o'chirish.

```go
err := os.Remove("data.txt")
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `os.Remove()` - faylni o'chirish.

### Directory Yaratish

**os.Mkdir()** - directory yaratish.

```go
err := os.Mkdir("mydir", 0755)
if err != nil {
    log.Fatal(err)
}
```

**Qoida:** `os.Mkdir()` - directory yaratish.

---

## 11. Muhim Qoidalar

### File I/O Best Practices

1. **defer file.Close()** - har doim faylni yopish
2. **Error handling** - har bir operatsiyani tekshirish
3. **Buffer** - katta fayllar uchun buffer ishlatish
4. **Permissions** - fayl huquqlarini to'g'ri belgilash

### JSON Best Practices

1. **JSON tags** - field nomlarini belgilash
2. **omitempty** - bo'sh fieldlarni o'tkazib yuborish
3. **Error handling** - JSON operatsiyalarini tekshirish
4. **Validation** - ma'lumotlarni tekshirish

---

## Xulosa

### File I/O

**Nima:** Fayllar bilan ishlash

**Vazifasi:**
- Fayl yaratish
- Fayl o'qish
- Fayl yozish
- Fayl o'chirish

**Afzalliklari:**
- Ma'lumotlarni saqlash
- Ma'lumotlarni almashish
- Configuration

### JSON

**Nima:** Ma'lumotlarni serializatsiya/deserializatsiya qilish

**Vazifasi:**
- Struct ni JSON ga o'tkazish
- JSON ni struct ga o'tkazish
- Ma'lumotlarni almashish

**Afzalliklari:**
- Oson format
- Ko'p tillar qo'llab-quvvatlaydi
- Web API uchun ideal

### Keyingi qadamlar

- Advanced file operations
- CSV, XML encoding/decoding
- File watching
- Real-world examples
