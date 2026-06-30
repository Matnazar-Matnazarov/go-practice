package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // pprof endpointlarini ro'yxatdan o'tkazish
)

// Yomon (optimallashtirilmagan) xotira ishlashi
func slowStringConcat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a" // Har bir tsiklda yangi xotira ajratiladi
	}
	return s
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := slowStringConcat(10000)
	fmt.Fprintf(w, "Uzunlik: %d", len(result))
}

func main() {
	http.HandleFunc("/concat", handler)

	fmt.Println("🚀 Pprof server ishga tushdi: http://localhost:6060")
	fmt.Println("Pprof manzillari:")
	fmt.Println(" - CPU Profile: http://localhost:6060/debug/pprof/profile?seconds=5")
	fmt.Println(" - Heap Profile: http://localhost:6060/debug/pprof/heap")

	// Serverni 6060 portda ishga tushiramiz
	log.Fatal(http.ListenAndServe(":6060", nil))
}
