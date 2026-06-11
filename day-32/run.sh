#!/bin/bash

echo "🚀 Benchmark natijalari:"
go test -bench=. -benchmem

echo ""
echo "🚀 Pprof Server ishga tushmoqda (CPU va xotira profillash uchun)..."
go run code.go
