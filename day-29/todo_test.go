package main

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	pb "day-29/todo"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestTodoService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	// Test Unary RPC (AddTodo)
	req := &pb.AddTodoRequest{Title: "Learn gRPC"}
	res, err := client.AddTodo(ctx, req)
	if err != nil {
		t.Fatalf("AddTodo failed: %v", err)
	}
	if res.Todo.Title != "Learn gRPC" {
		t.Errorf("Expected 'Learn gRPC', got '%v'", res.Todo.Title)
	}

	// Test Server Streaming RPC (ListTodos)
	stream, err := client.ListTodos(ctx, &pb.ListTodosRequest{})
	if err != nil {
		t.Fatalf("ListTodos failed: %v", err)
	}

	var todos []*pb.Todo
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Cannot receive: %v", err)
		}
		todos = append(todos, resp.Todo)
	}

	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}
}
