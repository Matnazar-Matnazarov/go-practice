package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Kun 24: CLI va Command-Line (flag, os.Args)")
	fmt.Println()

	// 1. os.Args
	fmt.Println("=== 1. os.Args ===")
	demonstrateOsArgs()
	fmt.Println()

	// 2. flag — parse qilingan flag'lar
	fmt.Println("=== 2. flag (Parse qilindi) ===")
	demonstrateFlags()
	fmt.Println()

	// 3. flag.Args() — qolgan argumentlar
	fmt.Println("=== 3. flag.Args() ===")
	demonstrateFlagArgs()
	fmt.Println()

	fmt.Println("=== Kun 24 yakunlandi! ===")
	fmt.Println("O'rganildi: os.Args, flag package, String/Int/Bool, Args va Usage")
}

func demonstrateOsArgs() {
	fmt.Printf("  os.Args[0] (dastur): %q\n", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Printf("  Birinchi argument: %q\n", os.Args[1])
	} else {
		fmt.Println("  Qo'shimcha argument yo'q. Misol: go run . arg1 arg2")
	}
}

func demonstrateFlags() {
	fmt.Printf("  port   = %d (usage: -port=8080)\n", *port)
	fmt.Printf("  verbose = %v (usage: -v yoki -verbose)\n", *verbose)
	fmt.Printf("  name  = %q (usage: -name=Ism)\n", *name)
}

func demonstrateFlagArgs() {
	args := flagArgs()
	if len(args) == 0 {
		fmt.Println("  flag.Args() bo'sh. Misol: go run . -port=80 file1.txt file2.txt")
		return
	}
	for i, a := range args {
		fmt.Printf("  flag.Args()[%d] = %q\n", i, a)
	}
}
