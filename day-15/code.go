package main

import "fmt"

func main() {
	fmt.Println("Kun 15: Packages & Modules (Advanced) - Internal, init order, import cycles")
	fmt.Println()

	fmt.Println("=== 1. Internal packages ===")
	fmt.Println()
	showInternalPackage()
	fmt.Println()

	fmt.Println("=== 2. init() order (concept) ===")
	fmt.Println()
	showInitOrderConcept()
	fmt.Println()

	fmt.Println("=== 3. Import cycles (concept) ===")
	fmt.Println()
	showImportCycleConcept()
	fmt.Println()

	fmt.Println("=== 4. go mod tidy / vendor (concept) ===")
	fmt.Println()
	showModCommandsConcept()
	fmt.Println()

	fmt.Println("=== Kun 15 yakunlandi! ===")
}
