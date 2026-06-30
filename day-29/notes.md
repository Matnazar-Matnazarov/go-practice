# Kun 29: gRPC va Protocol Buffers

## Kirish

Microservice'lar o'zaro qanday gaplashadi? REST API - keng tarqalgan, lekin JSON parsing va HTTP/1.1 sekin bo'lishi mumkin. **gRPC** (gRPC Remote Procedure Calls) - bu Google tomonidan yaratilgan, ochiq kodli va yuqori unumdorlikka ega RPC framework.

## 1. gRPC afzalliklari

- **Protocol Buffers (Protobuf)**: Ma'lumotlarni binar (binary) formatda serializatsiya qiladi. JSON'ga qaraganda kichikroq va tezroq.
- **HTTP/2**: Multiplexing, server push, header compression kabi imkoniyatlarni taqdim etadi. Bitta connection orqali ko'plab so'rovlar parallel bajariladi.
- **Streaming**: Unary (1:1), Server streaming (1:N), Client streaming (N:1), va Bidirectional streaming (N:N).
- **Code Generation**: Proto fayldan Go, Python, Java va boshqa tillar uchun server/client kodlarini avtomatik yaratadi.

## 2. Protocol Buffers (proto3)

API strukturasini `.proto` faylida yozamiz:

```proto
syntax = "proto3";
package todo;
option go_package = "day-29/todo";

service TodoService {
  rpc AddTodo (AddTodoRequest) returns (TodoResponse) {}
}

message AddTodoRequest {
  string title = 1; // 1 - bu maydonning (field) unikal tag raqami
}
```

## 3. Kodni generatsiya qilish

`protoc` va Go plugin'lari orqali `.proto` fayldan Go kodini yaratamiz:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    todo.proto
```
Bu `todo.pb.go` (ma'lumot turlari) va `todo_grpc.pb.go` (server/client interfeyslari) fayllarini yaratadi.

## 4. Serverni ishga tushirish

`TodoServiceServer` interfeysini implement qilamiz va serverni ro'yxatdan o'tkazamiz:

```go
type server struct {
    todo.UnimplementedTodoServiceServer
}

func (s *server) AddTodo(ctx context.Context, req *todo.AddTodoRequest) (*todo.TodoResponse, error) {
    // Logika...
    return &todo.TodoResponse{Todo: t}, nil
}
```

## 5. Test qilish (`bufconn`)

gRPC serverni tarmoq (TCP) portini ochmasdan, xotira (in-memory) orqali test qilish uchun `google.golang.org/grpc/test/bufconn` paketidan foydalaniladi. Bu unit testlarni juda tez va ishonchli qiladi.
