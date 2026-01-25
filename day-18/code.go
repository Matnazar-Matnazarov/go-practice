package main

import "fmt"

func main() {
	fmt.Println("Kun 18: Testing (Test Qilish)")
	fmt.Println()
	fmt.Println("Bu kunda Go testing asoslarini o'rganamiz.")
	fmt.Println("Test fayllarni ko'rish uchun quyidagi fayllarni tekshiring:")
	fmt.Println("  - math_test.go - Unit testing misollari")
	fmt.Println("  - table_test.go - Table-driven tests")
	fmt.Println("  - benchmark_test.go - Benchmark testing")
	fmt.Println("  - example_test.go - Example testing")
	fmt.Println()
	fmt.Println("Testlarni ishga tushirish:")
	fmt.Println("  go test                    # Barcha testlar")
	fmt.Println("  go test -v                 # Batafsil natija")
	fmt.Println("  go test -bench=.           # Benchmark testlar")
	fmt.Println("  go test -cover             # Coverage")
	fmt.Println()
	fmt.Println("=== Kun 18 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Test fayllari va test funksiyalari")
	fmt.Println("  ✓ Table-driven tests")
	fmt.Println("  ✓ Test helpers va setup/teardown")
	fmt.Println("  ✓ Error testing")
	fmt.Println("  ✓ Benchmark testing")
	fmt.Println("  ✓ Example testing")
	fmt.Println("  ✓ Test coverage")
}
