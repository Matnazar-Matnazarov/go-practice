package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketServer(t *testing.T) {
	// Hub'ni ishga tushirish
	hub := NewHub()
	go hub.Run()

	// HTTP Server handler'ni sozlash
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", hub.ServeHTTP)
	mux.HandleFunc("/clients", hub.GetClients)

	// Test server yaratish
	server := httptest.NewServer(mux)
	defer server.Close()

	// WebSocket URL yaratish
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

	// 1-Klient ulanishi
	conn1, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Klient 1 ulanmadi: %v", err)
	}
	defer conn1.Close()

	// Kichik pauza (hub ro'yxatga olishi uchun)
	time.Sleep(50 * time.Millisecond)

	// Klientlar sonini tekshirish (1 ta bo'lishi kerak)
	if count := hub.GetClientCount(); count != 1 {
		t.Errorf("Klientlar soni kutilganidek emas, kutilgan: 1, olingan: %d", count)
	}

	// 2-Klient ulanishi
	conn2, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Klient 2 ulanmadi: %v", err)
	}
	defer conn2.Close()

	time.Sleep(50 * time.Millisecond)
	if count := hub.GetClientCount(); count != 2 {
		t.Errorf("Klientlar soni kutilganidek emas, kutilgan: 2, olingan: %d", count)
	}

	// Broadcast orqali xabar yuborish
	testMessage := []byte("Salom, hammaga!")
	hub.Broadcast(testMessage)

	// Ikkala klient ham xabarni olishini tekshiramiz
	verifyMessageReceived(t, conn1, testMessage)
	verifyMessageReceived(t, conn2, testMessage)

	// Klient 1 ni uzish
	conn1.Close()
	time.Sleep(50 * time.Millisecond)

	if count := hub.GetClientCount(); count != 1 {
		t.Errorf("Klient uzilgandan so'ng count no'to'g'ri, kutilgan: 1, olingan: %d", count)
	}
}

func verifyMessageReceived(t *testing.T, conn *websocket.Conn, expected []byte) {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, msg, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Xabarni o'qishda xato: %v", err)
	}
	if string(msg) != string(expected) {
		t.Errorf("Noto'g'ri xabar. Kutilgan: %s, olingan: %s", expected, msg)
	}
}
