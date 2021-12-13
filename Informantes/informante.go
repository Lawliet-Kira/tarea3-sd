package main

import (
	"context"
	"fmt"
	"log"

	pb "lab3/game/helloworld"

	"google.golang.org/grpc"
)

const (
	BrokerAddress = "10.6.43.116:50051"
	defaultName   = "world"
)

var opcion = ""
var reloj_vector_Informante []int32

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
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println("Escoge a:")
	for opcion != "exit" {

		var opcion string

		// MENÚ
		fmt.Println("Escoge la opcion:")
		fmt.Println("	1. AddCity")
		fmt.Println("	2. UpdateName")
		fmt.Println("	3. UpdateNumbre")
		fmt.Println("	4. DeleteCity")

		fmt.Scanf("%s\n", &opcion)

		r, _ := client.Comands_Informantes_Broker(ctx, &pb.ComandIBRequest{Comand: opcion, RelojVector: reloj_vector_Informante})

		fmt.Println("Direccion IP seleccionada: ", r)
	}

}
