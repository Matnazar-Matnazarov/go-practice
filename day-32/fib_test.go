package main

import (
	"strings"
	"testing"
)

// Yomon usul: Stringlarni + bilan qo'shish
func concatStringsBad(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

// Yaxshi usul: strings.Builder ishlatish
func concatStringsGood(n int) string {
	var b strings.Builder
	b.Grow(n) // Xotirani oldindan ajratish
	for i := 0; i < n; i++ {
		b.WriteString("a")
	}
	return b.String()
}

// BenchmarkBad
func BenchmarkConcatBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringsBad(1000)
	}
}

// BenchmarkGood
func BenchmarkConcatGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringsGood(1000)
	}
}
