package main

import (
	"fmt"
)

func main() {
	fmt.Println("Kun 23: Database va SQL (database/sql)")
	fmt.Println()

	// 1. Databazaga ulanish
	fmt.Println("=== 1. Databazaga ulanish ===")
	demonstrateConnection()
	fmt.Println()

	// 2. CREATE TABLE
	fmt.Println("=== 2. CREATE TABLE ===")
	demonstrateCreateTable()
	fmt.Println()

	// 3. INSERT va Query
	fmt.Println("=== 3. INSERT va Query ===")
	demonstrateInsertAndQuery()
	fmt.Println()

	// 4. Prepared Statements
	fmt.Println("=== 4. Prepared Statements ===")
	demonstratePreparedStatements()
	fmt.Println()

	// 5. Transaction
	fmt.Println("=== 5. Transaction ===")
	demonstrateTransaction()
	fmt.Println()

	// 6. Context bilan ishlash
	fmt.Println("=== 6. Context bilan ishlash ===")
	demonstrateContext()
	fmt.Println()

	fmt.Println("=== Kun 23 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ sql.Open va connection pool")
	fmt.Println("  ✓ CREATE, INSERT, SELECT")
	fmt.Println("  ✓ Prepared statements (Prepare)")
	fmt.Println("  ✓ Transactions (Begin, Commit, Rollback)")
	fmt.Println("  ✓ Context bilan timeout")
}
