package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 13: Reflection (Refleksiya)")
	fmt.Println()

	// 1. REFLECTION ASOSLARI

	fmt.Println("=== 1. Reflection Asoslari ===")
	fmt.Println()

	// TypeOf va ValueOf
	fmt.Println("--- TypeOf va ValueOf ---")
	demonstrateTypeOfAndValueOf()
	fmt.Println()

	// Kind
	fmt.Println("--- Kind ---")
	demonstrateKind()
	fmt.Println()

	// 2. TYPE INSPECTION

	fmt.Println("=== 2. Type Inspection ===")
	fmt.Println()

	// Type name va kind
	fmt.Println("--- Type name va kind ---")
	demonstrateTypeNameAndKind()
	fmt.Println()

	// Struct fieldlar
	fmt.Println("--- Struct fieldlar ---")
	demonstrateStructFields()
	fmt.Println()

	// Field tags
	fmt.Println("--- Field tags ---")
	demonstrateFieldTags()
	fmt.Println()

	// Methodlar
	fmt.Println("--- Methodlar ---")
	demonstrateMethods()
	fmt.Println()

	// 3. VALUE MANIPULATION

	fmt.Println("=== 3. Value Manipulation ===")
	fmt.Println()

	// Value olish va o'zgartirish
	fmt.Println("--- Value olish va o'zgartirish ---")
	demonstrateValueGetAndSet()
	fmt.Println()

	// CanSet
	fmt.Println("--- CanSet ---")
	demonstrateCanSet()
	fmt.Println()

	// 4. STRUCT FIELDLAR BILAN ISHLASH

	fmt.Println("=== 4. Struct Fieldlar bilan Ishlash ===")
	fmt.Println()

	// Field qiymatlarini olish
	fmt.Println("--- Field qiymatlarini olish ---")
	demonstrateGetFieldValues()
	fmt.Println()

	// Field qiymatlarini o'zgartirish
	fmt.Println("--- Field qiymatlarini o'zgartirish ---")
	demonstrateSetFieldValues()
	fmt.Println()

	// FieldByName
	fmt.Println("--- FieldByName ---")
	demonstrateFieldByName()
	fmt.Println()

	// 5. METHODLAR BILAN ISHLASH

	fmt.Println("=== 5. Methodlar bilan Ishlash ===")
	fmt.Println()

	// Method chaqirish
	fmt.Println("--- Method chaqirish ---")
	demonstrateMethodCall()
	fmt.Println()

	// Method parametrlari
	fmt.Println("--- Method parametrlari ---")
	demonstrateMethodWithParameters()
	fmt.Println()

	// 6. SLICE VA MAP BILAN ISHLASH

	fmt.Println("=== 6. Slice va Map bilan Ishlash ===")
	fmt.Println()

	// Slice bilan ishlash
	fmt.Println("--- Slice bilan ishlash ---")
	demonstrateSliceReflection()
	fmt.Println()

	// Map bilan ishlash
	fmt.Println("--- Map bilan ishlash ---")
	demonstrateMapReflection()
	fmt.Println()

	// 7. TYPE ASSERTION VA TYPE SWITCH

	fmt.Println("=== 7. Type Assertion va Type Switch ===")
	fmt.Println()

	// Type assertion
	fmt.Println("--- Type assertion ---")
	demonstrateTypeAssertion()
	fmt.Println()

	// Type switch
	fmt.Println("--- Type switch ---")
	demonstrateTypeSwitch()
	fmt.Println()

	// 8. YAKUNIY XULOSA

	fmt.Println("=== Kun 13 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ reflect.Type va reflect.Value")
	fmt.Println("  ✓ Type inspection (name, kind, fields, methods)")
	fmt.Println("  ✓ Value manipulation (get, set, CanSet)")
	fmt.Println("  ✓ Struct fieldlar bilan ishlash")
	fmt.Println("  ✓ Methodlar bilan ishlash")
	fmt.Println("  ✓ Slice va Map bilan reflection")
	fmt.Println("  ✓ Type assertion va type switch")
}
