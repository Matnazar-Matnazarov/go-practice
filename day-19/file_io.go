package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)


// 1. FILE I/O ASOSLARI


// demonstrateFileOpenAndCreate - os.Open va os.Create
func demonstrateFileOpenAndCreate() {
	// os.Create - yangi fayl yaratish
	file, err := os.Create("test_output.txt")
	if err != nil {
		fmt.Printf("  Xato: %v\n", err)
		return
	}
	defer file.Close()

	// Faylga yozish
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Printf("  Yozish xatosi: %v\n", err)
		return
	}
	fmt.Println("  test_output.txt fayli yaratildi va yozildi")

	// os.Open - faylni o'qish
	readFile, err := os.Open("test_output.txt")
	if err != nil {
		fmt.Printf("  O'qish xatosi: %v\n", err)
		return
	}
	defer readFile.Close()

	// Fayldan o'qish
	data, err := io.ReadAll(readFile)
	if err != nil {
		fmt.Printf("  O'qish xatosi: %v\n", err)
		return
	}
	fmt.Printf("  O'qilgan ma'lumot: %s", string(data))
}

// demonstrateReadWriteFile - os.ReadFile va os.WriteFile
func demonstrateReadWriteFile() {
	// os.WriteFile - faylga yozish
	data := []byte("Bu os.WriteFile() orqali yozilgan ma'lumot\nIkkinchi qator\n")
	err := os.WriteFile("readwrite_test.txt", data, 0644)
	if err != nil {
		fmt.Printf("  Yozish xatosi: %v\n", err)
		return
	}
	fmt.Println("  readwrite_test.txt fayli yaratildi")

	// os.ReadFile - fayldan o'qish
	readData, err := os.ReadFile("readwrite_test.txt")
	if err != nil {
		fmt.Printf("  O'qish xatosi: %v\n", err)
		return
	}
	fmt.Printf("  O'qilgan ma'lumot:\n%s", string(readData))
}

// demonstrateBufioScanner - bufio.Scanner (qatorma-qator o'qish)
func demonstrateBufioScanner() {
	// Test fayl yaratish
	testData := "Birinchi qator\nIkkinchi qator\nUchinchi qator\n"
	err := os.WriteFile("scanner_test.txt", []byte(testData), 0644)
	if err != nil {
		fmt.Printf("  Fayl yaratish xatosi: %v\n", err)
		return
	}

	// Faylni ochish
	file, err := os.Open("scanner_test.txt")
	if err != nil {
		fmt.Printf("  Fayl ochish xatosi: %v\n", err)
		return
	}
	defer file.Close()

	// Scanner yaratish
	scanner := bufio.NewScanner(file)
	lineNum := 1
	fmt.Println("  Qatorma-qator o'qish:")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("    Qator %d: %s\n", lineNum, line)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("  Scanner xatosi: %v\n", err)
	}
}

// demonstrateBufioWriter - bufio.Writer (buffered yozish)
func demonstrateBufioWriter() {
	// Fayl yaratish
	file, err := os.Create("writer_test.txt")
	if err != nil {
		fmt.Printf("  Fayl yaratish xatosi: %v\n", err)
		return
	}
	defer file.Close()

	// Writer yaratish
	writer := bufio.NewWriter(file)
	
	// Yozish
	writer.WriteString("Buffered yozish - birinchi qator\n")
	writer.WriteString("Buffered yozish - ikkinchi qator\n")
	writer.WriteString("Buffered yozish - uchinchi qator\n")

	// Flush - buffer ni faylga yozish
	err = writer.Flush()
	if err != nil {
		fmt.Printf("  Flush xatosi: %v\n", err)
		return
	}
	fmt.Println("  writer_test.txt fayli buffered yozish orqali yaratildi")

	// O'qib tekshirish
	readData, err := os.ReadFile("writer_test.txt")
	if err != nil {
		fmt.Printf("  O'qish xatosi: %v\n", err)
		return
	}
	fmt.Printf("  O'qilgan ma'lumot:\n%s", string(readData))
}


// 3. FILE OPERATIONS


// demonstrateFileExists - File mavjudligini tekshirish
func demonstrateFileExists() {
	filename := "test_output.txt"
	
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("  %s fayli mavjud emas\n", filename)
		return
	}
	if err != nil {
		fmt.Printf("  Xato: %v\n", err)
		return
	}

	fmt.Printf("  %s fayli mavjud\n", filename)
	fmt.Printf("    O'lchami: %d bytes\n", info.Size())
	fmt.Printf("    Modifikatsiya vaqti: %v\n", info.ModTime())
	fmt.Printf("    Mode: %v\n", info.Mode())
}

// demonstrateFileRemove - File o'chirish
func demonstrateFileRemove() {
	// Test fayl yaratish
	testFile := "remove_test.txt"
	err := os.WriteFile(testFile, []byte("Bu fayl o'chiriladi"), 0644)
	if err != nil {
		fmt.Printf("  Fayl yaratish xatosi: %v\n", err)
		return
	}
	fmt.Printf("  %s fayli yaratildi\n", testFile)

	// Faylni o'chirish
	err = os.Remove(testFile)
	if err != nil {
		fmt.Printf("  Fayl o'chirish xatosi: %v\n", err)
		return
	}
	fmt.Printf("  %s fayli o'chirildi\n", testFile)

	// Tekshirish
	_, err = os.Stat(testFile)
	if os.IsNotExist(err) {
		fmt.Printf("  Tasdiqlandi: %s fayli mavjud emas\n", testFile)
	}
}

// demonstrateDirectoryCreate - Directory yaratish
func demonstrateDirectoryCreate() {
	// Directory yaratish
	dirName := "test_dir"
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("  %s directory allaqachon mavjud\n", dirName)
		} else {
			fmt.Printf("  Directory yaratish xatosi: %v\n", err)
			return
		}
	} else {
		fmt.Printf("  %s directory yaratildi\n", dirName)
	}

	// Directory ichida fayl yaratish
	testFile := filepath.Join(dirName, "test.txt")
	err = os.WriteFile(testFile, []byte("Directory ichidagi fayl"), 0644)
	if err != nil {
		fmt.Printf("  Fayl yaratish xatosi: %v\n", err)
		return
	}
	fmt.Printf("  %s fayli yaratildi\n", testFile)

	// Tozalash
	os.Remove(testFile)
	os.Remove(dirName)
	fmt.Printf("  Test directory va fayllar tozalandi\n")
}
