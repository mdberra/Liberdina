package main

import (
	"context"
	"fmt"
	"log"
	"net"

	usuariopb "github.com/Liberdina/protobuffers/seguridad/usuariopb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Get(ctx context.Context, req *usuariopb.UsuarioRequest) (*usuariopb.UsuarioResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	result := "Hello " + req.GetUsuario().GetName()

	res := &usuariopb.UsuarioResponse{
		Usuario: &usuariopb.Usuario{
			Name: result,
		},
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	usuariopb.RegisterUsuarioServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
