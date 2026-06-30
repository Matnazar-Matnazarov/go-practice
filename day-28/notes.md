# Kun 28: WebSocket va Real-Time Communication

## Kirish

HTTP — request/response modeliga asoslangan: klient so'rov yuboradi, server javob qaytaradi. Lekin real-time ilovalar (chat, live notifications, trading dashboard) uchun bu model yetarli emas — server klientga **o'zidan** ma'lumot yuborishi kerak.

**WebSocket** — ikki tomonlama (full-duplex) aloqa protokoli. Bir marta connection o'rnatilgandan keyin, server va klient bir-biriga istalgan vaqtda xabar yuborishi mumkin.

---

## 1. WebSocket nima?

### HTTP vs WebSocket

| Xususiyat | HTTP | WebSocket |
|-----------|------|-----------|
| **Model** | Request/Response | Full-duplex |
| **Connection** | Har so'rovda yangi (yoki Keep-Alive) | Doimiy (persistent) |
| **Server→Klient** | ❌ Mumkin emas | ✅ Mumkin |
| **Latency** | Yuqori (har so'rovda handshake) | Past (bir marta handshake) |
| **Use case** | REST API, CRUD | Chat, live updates, gaming |

### WebSocket handshake

1. Klient: `Upgrade: websocket` header bilan HTTP so'rov
2. Server: `101 Switching Protocols` javob
3. Connection TCP ga aylanadi — binary/text xabarlar oqimi

---

## 2. Gorilla WebSocket package

Go standart kutubxonasida WebSocket yo'q. **`gorilla/websocket`** — de-facto standart:

```bash
go get github.com/gorilla/websocket
```

### Conn (WebSocket connection)

```go
type Conn struct {
    // ...
}

// ReadMessage — xabar o'qish (blocking)
func (c *Conn) ReadMessage() (messageType int, p []byte, err error)

// WriteMessage — xabar yozish
func (c *Conn) WriteMessage(messageType int, data []byte) error

// Close — connection yopish
func (c *Conn) Close() error
```

### Message type

| Constant | Qiymat | Tavsif |
|----------|--------|--------|
| `TextMessage` | 1 | UTF-8 text |
| `BinaryMessage` | 2 | Binary data |
| `CloseMessage` | 8 | Connection close |
| `PingMessage` | 9 | Keep-alive |
| `PongMessage` | 10 | Keep-alive javob |

---

## 3. Upgrader (HTTP → WebSocket)

```go
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // Production'da CORS to'g'ri sozlash kerak:
    CheckOrigin: func(r *http.Request) bool {
        return true  // Demo uchun; production'da origin tekshirish!
    },
}
```

**Upgrade qilish:**

```go
func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("upgrade error: %v", err)
        return
    }
    defer conn.Close()
    
    // Connection bilan ishlash...
}
```

---

## 4. Connection management (Hub pattern)

Real loyihada ko'plab klientlar ulanadi. Har bir connection'ni alohida boshqarish qiyin. **Hub pattern** — barcha connection'larni markaziy boshqarish.

### Hub struct

```go
type Hub struct {
    clients    map[*websocket.Conn]bool  // Faol klientlar
    broadcast  chan []byte               // Xabar yuborish queue
    register   chan *websocket.Conn      // Yangi klient
    unregister chan *websocket.Conn      // Klient chiqdi
}
```

### Hub metodlari

```go
func (h *Hub) run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
            
        case client := <-h.unregister:
            delete(h.clients, client)
            client.Close()
            
        case message := <-h.broadcast:
            for client := range h.clients {
                client.WriteMessage(websocket.TextMessage, message)
            }
        }
    }
}
```

---

## 5. Client goroutine (Read/Write pump)

Har bir klient uchun 2 ta goroutine:

1. **Read pump** — klientdan xabar o'qish
2. **Write pump** — klientga xabar yuborish

### Read pump

```go
func (c *Client) readPump() {
    defer c.hub.unregister <- c.conn
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            break
        }
        // Broadcast qilish
        c.hub.broadcast <- message
    }
}
```

### Write pump

```go
func (c *Client) writePump() {
    for {
        select {
        case message := <-c.send:
            c.conn.WriteMessage(websocket.TextMessage, message)
        case <-c.done:
            return
        }
    }
}
```

---

## 6. Ping/Pong (Keep-alive)

WebSocket connection uzoq vaqt idle bo'lsa, firewall/router tomonidan yopilishi mumkin. **Ping/Pong** — connection tirikligini tekshirish.

```go
conn.SetReadDeadline(time.Now().Add(pongWait))
conn.SetPongHandler(func(string) error {
    conn.SetReadDeadline(time.Now().Add(pongWait))
    return nil
})

// Har 30 soniyada ping yuborish
ticker := time.NewTicker(pingPeriod)
for range ticker.C {
    conn.WriteMessage(websocket.PingMessage, nil)
}
```

---

## 7. Production best practices

### 1. CORS to'g'ri sozlash

```go
upgrader.CheckOrigin = func(r *http.Request) bool {
    origin := r.Header.Get("Origin")
    return origin == "https://myapp.com"
}
```

### 2. Rate limiting

Har bir klientdan kelayotgan xabarlar sonini cheklash (token bucket, sliding window).

### 3. Message validation

Klientdan kelgan xabaroldin tekshirish (JSON schema, max size).

### 4. Graceful shutdown

Server to'xtaganda barcha connection'larni to'g'ri yopish.

### 5. Error handling

Network xatolari, timeout, connection loss — hammasini log qilish va handle qilish.

---

## 8. Xulosa

**WebSocket** — real-time ilovalar uchun ikki tomonlama aloqa protokoli.

**O'rganildi:**
- HTTP vs WebSocket farqi
- Gorilla WebSocket package
- Upgrader (HTTP → WebSocket)
- Hub pattern (connection management)
- Read/Write pump (client goroutines)
- Ping/Pong (keep-alive)

**Keyingi qadamlar:**
- JWT authentication WebSocket bilan
- Message persistence (database)
- Horizontal scaling (Redis pub/sub)
- Binary message support

---

## 9. Demo ishga tushirish

```bash
cd day-28
./run.sh
```

Keyin brauzerda `http://localhost:8080` ochib, chat test qilish mumkin.
