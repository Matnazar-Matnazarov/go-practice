package main

import "fmt"

// checkAge - if/else shartli operatori
func checkAge(age int) {
	if age >= 18 {
		fmt.Printf("Yosh: %d - Kattalar uchun\n", age)
	} else {
		fmt.Printf("Yosh: %d - Yosh bolalar uchun\n", age)
	}
}

// checkGrade - if/else if/else zanjiri
func checkGrade(score int) string {
	if score >= 90 {
		return "A (A'lo)"
	} else if score >= 80 {
		return "B (Yaxshi)"
	} else if score >= 70 {
		return "C (Qoniqarli)"
	} else if score >= 60 {
		return "D (Qoniqarsiz)"
	} else {
		return "F (Yiqilgan)"
	}
}

// checkNumber - if bilan murakkab shartlar
func checkNumber(num int) {
	if num > 0 {
		fmt.Printf("%d - musbat son\n", num)
	} else if num < 0 {
		fmt.Printf("%d - manfiy son\n", num)
	} else {
		fmt.Printf("%d - nol\n", num)
	}
}

// checkEvenOdd - if bilan qoldiq tekshirish
func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d - juft son\n", num)
	} else {
		fmt.Printf("%d - toq son\n", num)
	}
}

// checkRange - if bilan oraliq tekshirish
func checkRange(num int) {
	if num >= 1 && num <= 10 {
		fmt.Printf("%d - 1 va 10 orasida\n", num)
	} else if num > 10 && num <= 20 {
		fmt.Printf("%d - 11 va 20 orasida\n", num)
	} else {
		fmt.Printf("%d - oraliqdan tashqarida\n", num)
	}
}

// checkMultipleConditions - bir nechta shartlar
func checkMultipleConditions(age int, hasLicense bool) {
	if age >= 18 && hasLicense {
		fmt.Println("Haydashga ruxsat beriladi")
	} else if age >= 18 && !hasLicense {
		fmt.Println("Yosh yetarli, lekin haydovchilik guvohnomasi yo'q")
	} else {
		fmt.Println("Haydashga ruxsat berilmaydi")
	}
}

// checkDivisibility - if bilan bo'linish tekshirish
func checkDivisibility(num, divisor int) {
	if divisor == 0 {
		fmt.Println("Nolga bo'lish mumkin emas!")
		return
	}
	if num%divisor == 0 {
		fmt.Printf("%d soni %d ga bo'linadi\n", num, divisor)
	} else {
		fmt.Printf("%d soni %d ga bo'linmaydi\n", num, divisor)
	}
}

// switchBasic - oddiy switch operatori
func switchBasic(day int) {
	switch day {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Noto'g'ri kun raqami")
	}
}

// switchMultiple - bir nechta case qiymatlari
func switchMultiple(month int) {
	switch month {
	case 12, 1, 2:
		fmt.Println("Qish")
	case 3, 4, 5:
		fmt.Println("Bahor")
	case 6, 7, 8:
		fmt.Println("Yoz")
	case 9, 10, 11:
		fmt.Println("Kuz")
	default:
		fmt.Println("Noto'g'ri oy raqami")
	}
}

// switchExpression - switch expression bilan
func switchExpression(score int) string {
	switch {
	case score >= 90:
		return "A'lo"
	case score >= 80:
		return "Yaxshi"
	case score >= 70:
		return "Qoniqarli"
	case score >= 60:
		return "Qoniqarsiz"
	default:
		return "Yiqilgan"
	}
}

// switchType - type switch (keyingi kunlarda batafsil)
func switchType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Noma'lum tip: %T\n", v)
	}
}

// checkLogin - amaliy misol: login tekshirish
func checkLogin(username, password string) bool {
	validUsername := "admin"
	validPassword := "12345"

	if username == validUsername && password == validPassword {
		return true
	}
	return false
}

// checkTemperature - amaliy misol: harorat tekshirish
func checkTemperature(temp float64) string {
	switch {
	case temp < 0:
		return "Muzlash"
	case temp >= 0 && temp < 15:
		return "Sovuq"
	case temp >= 15 && temp < 25:
		return "Iqlim"
	case temp >= 25 && temp < 35:
		return "Issiq"
	default:
		return "Juda issiq"
	}
}

// checkLeapYear - amaliy misol: kabisa yili tekshirish
func checkLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

// checkVotingEligibility - amaliy misol: saylov huquqi
func checkVotingEligibility(age int, isCitizen bool) string {
	if age >= 18 && isCitizen {
		return "Saylovda qatnashish mumkin"
	} else if age < 18 {
		return "Yosh yetarli emas"
	} else if !isCitizen {
		return "Fuqarolik yo'q"
	}
	return "Saylovda qatnashish mumkin emas"
}
