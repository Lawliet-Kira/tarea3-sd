package main

import (
	"context"
	pb "lab3/game/helloworld"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComunicationServer
}

func (s *server) Comands_Informantes_Broker(ctx context.Context, in *pb.ComandIBRequest) (*pb.ComandIBReply, error) {
	rand.Seed(time.Now().UnixNano())
	log.Printf("Operacion Received: %v", in.GetOperacion())
	reloj_vector_Informante := in.GetRelojVector()
	var ip = ""

	if len(reloj_vector_Informante) == 0 {

		Rand_num := rand.Intn(3)

		if Rand_num == 0 {
			ip = Server1Address
		}
		if Rand_num == 1 {
			ip = Server2Address
		}
		if Rand_num == 2 {
			ip = Server3Address
		}

	} else {
		//AQUÍ HAY QUE ENVIAR MENSAJE A FULCRUMS TAL QUE ME RETORNE SUS RELOJES DE VECTORES

		/* CODIGO */

		/*reloj_vector_s1 := rvs1
		reloj_vector_s2 := rvs2
		reloj_vector_s3 := rvs3*/

		//AQUÍ SE COMPARAN LOS RELOJES PARA VER CUAL SERVER ES EL QUE CUMPLE CON "READ YOUR WRITES

		/*CODIGO*/

	}

	//DEBUG

	ip = "10.6.43.114:50052"

	return &pb.ComandIBReply{Ip: ip}, nil
}

func (s *server) Comands_Leia_Broker(ctx context.Context, in *pb.ComandLBRequest) (*pb.ComandLBReply, error) {
	rand.Seed(time.Now().UnixNano())
	log.Printf("Operacion Received: %v", in.GetOperacion())
	reloj_vector_Leia := in.GetRelojVector()
	var ip = ""

	if len(reloj_vector_Informante) == 0 {

		Rand_num := rand.Intn(3)

		if Rand_num == 0 {
			ip = Server1Address
		}
		if Rand_num == 1 {
			ip = Server2Address
		}
		if Rand_num == 2 {
			ip = Server3Address
		}

	} else {
		/*CODE*/
	}

	// REALIZAR CONEXION

	return &pb.ComandIBReply{Ip: ip}, nil
}

const (
	port           = ":50051"
	BrokenAddress  = "10.6.43.116:50052"
	Server1Address = "10.6.43.113:50052"
	Server2Address = "10.6.43.114:50052"
	Server3Address = "10.6.43.115:50052"
)

func main() {
	// Crear un gRPC canal para comunicarse con el servidor
	// 	-> Esto se crea pasando server address y port number a grpc.Dial()

	// Puerto para escucha de los clientes
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	// Crea una isntancia de server gRPC
	s := grpc.NewServer()

	// Registrar nuestra implementación de server con gRPC server
	pb.RegisterComunicationServer(s, &server{})

	log.Printf("Server Broker escuchando en %v", lis.Addr())

	// Llamar Server() con los detalles de puerto para realizar un bloquero
	// de espera hasta que el proceso sea killed o Stop() es llamado
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
