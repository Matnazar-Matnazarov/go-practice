package main

import (
	"context"
	"fmt"
	pb "day-29/todo"
)

type server struct {
	pb.UnimplementedTodoServiceServer
	todos []*pb.Todo
}

func (s *server) AddTodo(ctx context.Context, req *pb.AddTodoRequest) (*pb.TodoResponse, error) {
	todo := &pb.Todo{
		Id:        fmt.Sprintf("%d", len(s.todos)+1),
		Title:     req.Title,
		Completed: false,
	}
	s.todos = append(s.todos, todo)
	return &pb.TodoResponse{Todo: todo}, nil
}

func (s *server) ListTodos(req *pb.ListTodosRequest, stream pb.TodoService_ListTodosServer) error {
	for _, todo := range s.todos {
		if err := stream.Send(&pb.TodoResponse{Todo: todo}); err != nil {
			return err
		}
	}
	return nil
}
