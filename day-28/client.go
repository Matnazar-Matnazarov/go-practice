package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Write vaqti timeout
	writeWait = 10 * time.Second

	// Pong kutish vaqti
	pongWait = 60 * time.Second

	// Ping period (pongWait dan kichik bo'lishi kerak)
	pingPeriod = (pongWait * 9) / 10

	// Max message size
	maxMessageSize = 512
)

// Client bitta WebSocket connection'ni ifodalaydi.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// ReadPump klientdan xabar o'qiydi va hub'ga yuboradi.
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("websocket error: %v", err)
			}
			break
		}
		// Broadcast message
		c.hub.Broadcast(message)
	}
}

// WritePump hub'dan xabar o'qiydi va klientga yuboradi.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub closed channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Queue'dagi barcha xabarlarni qo'shish
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeHTTP WebSocket connection'ni qabul qiladi.
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:  h,
		conn: conn,
		send: make(chan []byte, 256),
	}

	client.hub.register <- client

	// Read va Write pump'larni goroutine'da ishga tushirish
	go client.WritePump()
	go client.ReadPump()
}

// BroadcastMessage barcha klientlarga xabar yuboradi (HTTP endpoint uchun).
func (h *Hub) BroadcastMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	defer r.Body.Close()

	message := buf.Bytes()
	if len(message) == 0 {
		http.Error(w, "Empty message", http.StatusBadRequest)
		return
	}

	h.Broadcast(message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message broadcasted"))
}

// GetClients faol klientlar sonini qaytaradi (HTTP endpoint).
func (h *Hub) GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count := h.GetClientCount()
	w.Write([]byte(`{"clients":` + strconv.Itoa(count) + `}`))
}
