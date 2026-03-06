# Kun 23: Database va SQL (database/sql)

## Kirish

Kun 20–22 da HTTP server va clientni o‘rgandik. Real loyihalarda esa serverlar ma’lumotlarni **database** da saqlaydi. Go’da `database/sql` package — SQL databazalar (SQLite, PostgreSQL, MySQL) bilan ishlashning standard va production-ready yechimi.

---

## 1. database/sql nima?

**database/sql** — SQL databazalarga ulanish, query yuborish va natijalarni o‘qish uchun standart interface.

### Qaysi databazalar bilan ishlaydi?

- **SQLite** — yengil, file-based (localhost/embedded)
- **PostgreSQL** — production darajadagi open-source
- **MySQL/MariaDB** — eng mashhur
- **SQL Server, Oracle** — enterprise

### Asosiy tushunchalar

| Komponent | Vazifasi |
|-----------|----------|
| `sql.DB` | Connection pool — bir nechta ulanish boshqaruvi |
| `sql.Tx` | Transaction — atomik operatsiyalar |
| `sql.Stmt` | Prepared statement — parametrlangan query |
| `sql.Rows` | Query natijalari (iterator) |
| `sql.Row` | Bir qator natija |

---

## 2. Databazaga ulanish

### sql.Open — driver va DSN

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"  // Driver import
)

// SQLite in-memory
db, err := sql.Open("sqlite3", ":memory:")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

**Muhim:** `sql.Open` ulanish emas, faqat pool yaratadi. Haqiqiy ulanish `db.Ping()` bilan tekshiriladi.

### Connection Pool sozlamalari

```go
db.SetMaxOpenConns(25)       // Maksimal ochiq ulanish
db.SetMaxIdleConns(10)       // Bo‘sh ulanishlar soni
db.SetConnMaxLifetime(5*time.Minute) // Ulanish umri
```

**Best practice:** Pool sozlamalari production’da shart — leak va oshib ketishni oldini oladi.

---

## 3. CREATE, INSERT, SELECT

### Exec — INSERT, UPDATE, DELETE

```go
result, err := db.Exec(
    "INSERT INTO users (name, email) VALUES (?, ?)",
    "Ali", "ali@example.com",
)
if err != nil {
    log.Fatal(err)
}

// Natija ma’lumotlari
lastID, _ := result.LastInsertId()
rowsAffected, _ := result.RowsAffected()
```

**Placeholder:** `?` (SQLite, MySQL) yoki `$1, $2` (PostgreSQL).

### Query — SELECT (ko‘p qator)

```go
rows, err := db.Query("SELECT id, name FROM users WHERE active = ?", true)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()  // Muhim!

for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil {
        log.Fatal(err)
    }
    fmt.Println(id, name)
}
```

### QueryRow — SELECT (bir qator)

```go
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
if err == sql.ErrNoRows {
    fmt.Println("Foydalanuvchi topilmadi")
}
```

---

## 4. Prepared Statements

Prepared statement — bir marta tayyorlab, ko‘p marta ishlatish. SQL injection’ga qarshi himoya va performance uchun.

```go
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Ko‘p marta insert
for _, u := range users {
    _, err := stmt.Exec(u.Name, u.Email)
    if err != nil {
        log.Fatal(err)
    }
}
```

**Afzallik:**
- SQL injectiondan xavfsiz
- Plan cache (tezroq bajarish)
- Bir nechta insert/update uchun optimal

---

## 5. Transaction

Transaction — bir nechta operatsiyani atomik bajarish. Hammasi muvaffaqiyatli bo‘lsa `Commit()`, aks holda `Rollback()`.

```go
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}

// Transfer: A dan B ga pul o‘tkazish
_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", 100, 1)
if err != nil {
    tx.Rollback()
    return
}

_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", 100, 2)
if err != nil {
    tx.Rollback()
    return
}

// Hammasi muvaffaqiyatli
if err := tx.Commit(); err != nil {
    log.Fatal(err)
}
```

**Qo‘llanish joylari:**
- Bank transferlari
- Order + inventory update
- Multi-table insertlar

---

## 6. Context bilan ishlash

Context — query’ga timeout va cancellation qo‘shish.

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", 1)
var name string
if err := row.Scan(&name); err != nil {
    // context.DeadlineExceeded — timeout
}
```

**Muhim method’lar:**
- `QueryContext`, `QueryRowContext`
- `ExecContext`
- `BeginTx` — transaction + context

---

## 7. Best Practices

### Xavfsizlik

1. **Har doim placeholder ishlating** — `?` yoki `$1`. String concatenation qilmang!
2. **sql.NullString/NullInt64** — NULL qiymatlar uchun
3. **defer rows.Close()** — resource leak oldini olish

### Performance

1. **Connection pool** — default cheksiz emas, sozlang
2. **Prepared statements** — ko‘p takrorlanuvchi query uchun
3. **Transaction** — bir nechta operatsiyani birlashtiring
4. **Context timeout** — uzun query’larni cheklang

### Error handling

```go
if err == sql.ErrNoRows {
    // Ma’lumot yo‘q — normal holat
} else if err != nil {
    // Boshqa xato — log qilish
}
```

---

## Xulosa

### O‘rganildi:

- `sql.Open` va connection pool
- `Exec`, `Query`, `QueryRow`
- Prepared statements (`Prepare`, `Stmt`)
- Transaction (`Begin`, `Commit`, `Rollback`)
- Context bilan timeout

### Production uchun:

- **sqlc** — SQL’dan Go code generate qilish
- **GORM** — ORM (agar kerak bo‘lsa)
- **pgx** — PostgreSQL uchun tezroq driver

---

## Keyingi qadamlar

- Day 24: gRPC yoki WebSocket
- Day 25: Docker bilan deploy
- Day 26: CI/CD (GitHub Actions)
