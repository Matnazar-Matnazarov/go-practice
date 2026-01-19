package main

import "fmt"

func showInternalPackage() {
	fmt.Println("  internal/ papkasi ichidagi package faqat ota (parent) package'lar tomonidan import qilinadi")
	fmt.Println("  Misol: project/internal/logger -> project/cmd/app import qila oladi")
	fmt.Println("  Lekin project/x/logger import qila olmaydi")
}

func showInitOrderConcept() {
	fmt.Println("  init() avtomatik chaqiriladi:")
	fmt.Println("    1) import qilingan package'larning init()")
	fmt.Println("    2) joriy package init()")
	fmt.Println("    3) main()")
	fmt.Println("  Bir package ichida bir nechta init() bo'lishi mumkin")
}

func showImportCycleConcept() {
	fmt.Println("  Import cycle - A -> B -> A kabi aylanma bog'liqlik")
	fmt.Println("  Go bunday holatda build paytida xato beradi")
	fmt.Println("  Yechimlar: umumiy kodni 3-chi package'ga ko'chirish, interfeyslar bilan ajratish")
}

func showModCommandsConcept() {
	fmt.Println("  go mod tidy  - ishlatilmayotgan dependency'larni tozalaydi, keraklilarini qo'shadi")
	fmt.Println("  go mod vendor - dependency'larni vendor/ ichiga nusxalaydi")
	fmt.Println("  go list -m all - barcha modullar ro'yxati")
}
