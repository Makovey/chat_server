package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	m "github.com/Makovey/chat-server/pkg/chat/v1"
)

const grpcPort = 3001

type server struct {
	m.UnimplementedChatV1Server
}

func (s *server) Create(context.Context, *m.CreateRequest) (*m.CreateResponse, error) {
	return &m.CreateResponse{Id: 3001}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	m.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
