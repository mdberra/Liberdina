package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Liberdina/protobuffers/seguridad/usuariopb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := usuariopb.NewUsuarioServiceClient(cc)

	doUnary(c)
}

func doUnary(c usuariopb.UsuarioServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &usuariopb.UsuarioRequest{
		Usuario: &usuariopb.Usuario{
			Name: "Stephane",
		},
	}
	res, err := c.Get(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Usuario)
}
