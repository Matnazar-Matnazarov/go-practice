package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// demonstrateConnection — databazaga ulanish va connection pool
type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

func demonstrateConnection() {
	// SQLite in-memory databaza
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Printf("Open error: %v", err)
		return
	}
	defer db.Close()

	// Connection pool sozlamalari
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Printf("Ping error: %v", err)
		return
	}

	fmt.Println("  ✓ sql.Open() — driver va DSN bilan")
	fmt.Println("  ✓ db.Ping() — ulanishni tekshirish")
	fmt.Println("  ✓ Connection pool sozlamalari")
}

func demonstrateCreateTable() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Printf("Open error: %v", err)
		return
	}
	defer db.Close()

	createSQL := `CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	if _, err := db.Exec(createSQL); err != nil {
		log.Printf("Create table error: %v", err)
		return
	}

	fmt.Println("  ✓ db.Exec() — CREATE TABLE")
	fmt.Println("  ✓ SQL schema yaratildi")
}

func demonstrateInsertAndQuery() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Printf("Open error: %v", err)
		return
	}
	defer db.Close()

	// Table yaratish
	db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		created_at DATETIME
	)`)

	// INSERT
	result, err := db.Exec(
		"INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)",
		"Ali", "ali@example.com", time.Now(),
	)
	if err != nil {
		log.Printf("Insert error: %v", err)
		return
	}
	lastID, _ := result.LastInsertId()
	fmt.Printf("  ✓ INSERT: last ID = %d\n", lastID)

	// SELECT
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Printf("Query error: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		fmt.Printf("  ✓ SELECT: %+v\n", u)
	}
}

func demonstratePreparedStatements() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Printf("Open error: %v", err)
		return
	}
	defer db.Close()

	db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)`)

	// Prepared statement yaratish
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		log.Printf("Prepare error: %v", err)
		return
	}
	defer stmt.Close()

	// Bir nechta insert
	users := []struct {
		name  string
		email string
	}{
		{"Ali", "ali@test.com"},
		{"Vali", "vali@test.com"},
		{"Guli", "guli@test.com"},
	}

	for _, u := range users {
		if _, err := stmt.Exec(u.name, u.email); err != nil {
			log.Printf("Exec error: %v", err)
			continue
		}
	}

	fmt.Println("  ✓ db.Prepare() — prepared statement")
	fmt.Printf("  ✓ Bir statement bilan %d ta insert\n", len(users))
}

func demonstrateTransaction() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Printf("Open error: %v", err)
		return
	}
	defer db.Close()

	db.Exec(`CREATE TABLE accounts (id INTEGER PRIMARY KEY, balance INTEGER)`)
	db.Exec("INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 500)")

	// Transaction boshlash
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Begin error: %v", err)
		return
	}

	// Transfer: 1 dan 2 ga 200
	if _, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", 200, 1); err != nil {
		tx.Rollback()
		log.Printf("Transaction rollback: %v", err)
		return
	}

	if _, err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", 200, 2); err != nil {
		tx.Rollback()
		log.Printf("Transaction rollback: %v", err)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Commit error: %v", err)
		return
	}

	fmt.Println("  ✓ db.Begin() — transaction boshlash")
	fmt.Println("  ✓ tx.Commit() — saqlash")
	fmt.Println("  ✓ tx.Rollback() — bekor qilish")
	fmt.Println("  ✓ Atomic transfer misoli")
}

func demonstrateContext() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Printf("Open error: %v", err)
		return
	}
	defer db.Close()

	db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)`)
	db.Exec("INSERT INTO users VALUES (1, 'Ali')")

	// Context bilan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", 1)
	var name string
	if err := row.Scan(&name); err != nil {
		log.Printf("Scan error: %v", err)
		return
	}

	fmt.Println("  ✓ db.QueryRowContext() — context bilan")
	fmt.Printf("  ✓ Natija: name=%s\n", name)
}
