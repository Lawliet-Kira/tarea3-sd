package main

import (
	"bufio"
	"context"
	"fmt"
	pb "lab3/game/helloworld"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

const (
	BrokerAddress = "10.6.43.116:50051"
	defaultName   = "world"
)

var opcion = ""
var reloj_vector_Leia []int32

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

	for opcion != "exit" {

		var operacion string
		var valor string
		var comand string

		// MENÚ
		fmt.Println("Ingrese Operación: ")
		fmt.Println("	1. Obtener rebeldes")

		fmt.Scanf("%s\n", &operacion)

		// clean the prompt
		CallClear()

		fmt.Println("\nIngresar comando: ")

		// El escáner puede escanear entradas por líneas

		inputReader := bufio.NewReader(os.Stdin)
		comand, _ = inputReader.ReadString('\n')
		fmt.Println(comand)

		splited := strings.Split(comand, " ")

		nombre_planeta := splited[0]
		nombre_ciudad := splited[1]

		fmt.Println("nombrePlaneta: ", nombre_planeta)
		fmt.Println("nombreCiudad: ", nombre_ciudad)

		r, _ := client.Comands_Leia_Broker(ctx, &pb.ComandLBRequest{Operacion: operacion, NombrePlaneta: nombre_planeta, NombreCiudad: nombre_ciudad})

		cant_soldados := r.GetCantRebelds()

		reloj_vector_Leia := r.GetRelojVector()

		fmt.Println("Cantidad de soldados: ", cant_soldados)
		fmt.Println("Reloj vector Leia: ", reloj_vector_Leia)

	}

}
