#!/bin/bash

# Kun 28: WebSocket Demo
# Bu skript WebSocket serverni ishga tushiradi

set -e

echo "🚀 Kun 28: WebSocket va Real-Time Communication"
echo "================================================"
echo ""

# Dependencies o'rnatish
echo "📦 Dependencies tekshirilmoqda..."
if [ ! -f "go.mod" ]; then
    echo "📝 go.mod yaratilmoqda..."
    go mod init day-28
fi

echo "📥 gorilla/websocket o'rnatilmoqda..."
go get github.com/gorilla/websocket

echo ""
echo "⚙️  Server ishga tushirilmoqda..."
echo "   - WebSocket: ws://localhost:8080/ws"
echo "   - Frontend:  http://localhost:8080"
echo "   - API:       POST http://localhost:8080/broadcast"
echo ""
echo "💡 Maslahat: Bir nechta browser tab ochib, real-time chat test qiling!"
echo ""
echo "🛑 To'xtatish: Ctrl+C"
echo "================================================"
echo ""

go run *.go
