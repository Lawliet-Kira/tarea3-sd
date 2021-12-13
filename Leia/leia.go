package main

import (
	"context"
	pb "lab3/game/helloworld"
	"log"

	"google.golang.org/grpc"
)

const (
	BrokerAddress = "10.6.43.116:50051"
	defaultName   = "world"
)

func main() {
	// Crear un gRPC canal para comunicarse con el servidor
	// 	-> Esto se crea pasando server address y port number a grpc.Dial()
	conn, err := grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Client Stub to perform RPCs
	client := pb.NewComunicationClient(conn)
	// Contact the server and psirint out its response.
	ctx := context.Background()

}
