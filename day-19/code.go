package main

import "fmt"

func main() {
	fmt.Println("Kun 19: File I/O va JSON (Fayl Operatsiyalari va JSON)")
	fmt.Println()

	// 1. FILE I/O ASOSLARI

	fmt.Println("=== 1. File I/O Asoslari ===")
	fmt.Println()

	// os.Open va os.Create
	fmt.Println("--- os.Open va os.Create ---")
	demonstrateFileOpenAndCreate()
	fmt.Println()

	// os.ReadFile va os.WriteFile
	fmt.Println("--- os.ReadFile va os.WriteFile ---")
	demonstrateReadWriteFile()
	fmt.Println()

	// bufio.Scanner
	fmt.Println("--- bufio.Scanner (qatorma-qator o'qish) ---")
	demonstrateBufioScanner()
	fmt.Println()

	// bufio.Writer
	fmt.Println("--- bufio.Writer (buffered yozish) ---")
	demonstrateBufioWriter()
	fmt.Println()

	// 2. JSON ENCODING/DECODING

	fmt.Println("=== 2. JSON Encoding/Decoding ===")
	fmt.Println()

	// json.Marshal va json.Unmarshal
	fmt.Println("--- json.Marshal va json.Unmarshal ---")
	demonstrateJSONMarshalUnmarshal()
	fmt.Println()

	// json.MarshalIndent
	fmt.Println("--- json.MarshalIndent (formatlangan JSON) ---")
	demonstrateJSONMarshalIndent()
	fmt.Println()

	// json.Encoder va json.Decoder
	fmt.Println("--- json.Encoder va json.Decoder (stream) ---")
	demonstrateJSONEncoderDecoder()
	fmt.Println()

	// JSON tags
	fmt.Println("--- JSON tags ---")
	demonstrateJSONTags()
	fmt.Println()

	// 3. FILE OPERATIONS

	fmt.Println("=== 3. File Operations ===")
	fmt.Println()

	// File mavjudligini tekshirish
	fmt.Println("--- File mavjudligini tekshirish ---")
	demonstrateFileExists()
	fmt.Println()

	// File o'chirish
	fmt.Println("--- File o'chirish ---")
	demonstrateFileRemove()
	fmt.Println()

	// Directory yaratish
	fmt.Println("--- Directory yaratish ---")
	demonstrateDirectoryCreate()
	fmt.Println()

	// 4. YAKUNIY XULOSA

	fmt.Println("=== Kun 19 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ os.Open, os.Create, os.OpenFile")
	fmt.Println("  ✓ os.ReadFile va os.WriteFile")
	fmt.Println("  ✓ bufio.Scanner va bufio.Writer")
	fmt.Println("  ✓ json.Marshal va json.Unmarshal")
	fmt.Println("  ✓ json.MarshalIndent")
	fmt.Println("  ✓ json.Encoder va json.Decoder")
	fmt.Println("  ✓ JSON tags va field options")
	fmt.Println("  ✓ File operations (exists, remove, directory)")
}
