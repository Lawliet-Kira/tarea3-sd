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

	// reloj_vector_Informante := in.GetRelojVector()
	var ip = ""
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
	//Aqui pondria mi Read Your Writes, si tan solo lo hubiera programado

	return &pb.ComandIBReply{Ip: ip}, nil
}

type Keyvalue struct {
	planeta string
	vector  []int32
}

func EscogerIP_MonotonicReads(reloj_s1 []int32, reloj_s2 []int32, reloj_s3 []int32, reloj_Leia []int32) string{
	
	var ip_seleccionadas []string
	
	flag := true
	
	for i, pos range := reloj_Leia {
		if reloj_s1[i] < pos{
			flag = false
			break
		}
	}

	if flag {
		ip_seleccionadas = append(ip_seleccionadas, ServerAddress1)
	}
	
	flag = true
	
	for i, pos range := reloj_Leia {
		if reloj_s2[i] < pos{
			flag = false
			break
		}
		
	}
	
	if flag {
		ip_seleccionadas = append(ip_seleccionadas, ServerAddress2)
	}
	flag = true
	for i, pos range := reloj_Leia{
		if reloj_s3[i] < pos{
			flag = false
			break
		}
	}
	
	if flag {
		ip_seleccionadas = append(ip_seleccionadas, ServerAddress3)
	}
	rand.Seed(time.Now().UnixNano())
	Rand_num := rand.Intn(len(ip_seleccionadas))
	OutputIp := ip_seleccionadas[Rand_num]
	
	return OutputIp

}

func (s *server) Comands_Leia_Broker(ctx context.Context, in *pb.ComandLBRequest) (*pb.ComandLBReply, error) {

	rand.Seed(time.Now().UnixNano())
	log.Printf("Operacion Received: %v", in.GetOperacion())
	reloj_vector_Leia := in.GetRelojVector()
	var ip = ""
	var cant_soldados = "0"
	operacion := in.GetOperacion()
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()

	// Pedir tabla hash de cada server

	//SERVER 1

	conn, err := grpc.Dial(Server1Address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Client Stub to perform RPCs
	client := pb.NewComunicationClient(conn)
	// Contact the server and psirint out its response.
	defer conn.Close()
	ctx := context.Background()
	var reloj_vector_s1 []int32
	var reloj_vector_s2 []int32
	var reloj_vector_s3 []int32
	r, _ := client.Comands_Request_Hashing(ctx, &pb.PingMsg{Signal: signal})
	hash_s1 := r.GetHashing()
	indice_planeta_s1 := findHashing(hash_s1)
	if indice_planeta_s1 == -1{
		reloj_vector_s1 = {0,0,0}
	}else{
		reloj_vector_s1 = hash_s1[indice_planeta_s1].vector
	}
	
	
	//SERVER 2

	conn, err = grpc.Dial(Server2Address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Client Stub to perform RPCs
	client = pb.NewComunicationClient(conn)
	// Contact the server and print out its response.
	ctx = context.Background()

	r, _ = client.Comands_Request_Hashing(ctx, &pb.PingMsg{Signal: signal})
	hash_s2 := r.GetHashing()
	indice_planeta_s2 := findHashing(hash_s2)
	if indice_planeta_s2 == -1{
		reloj_vector_s2 = {0,0,0}
	}else{
		reloj_vector_s2 = hash_s1[indice_planeta_s2].vector
	}

	defer conn.Close()

	//SERVER 3

	conn, err = grpc.Dial(Server3Address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Client Stub to perform RPCs
	client = pb.NewComunicationClient(conn)
	// Contact the server and psirint out its response.
	ctx = context.Background()

	r, _ = client.Comands_Request_Hashing(ctx, &pb.PingMsg{Signal: signal})
	hash_s3 := r.GetHashing()
	indice_planeta_s3 := findHashing(hash_s3)
	if indice_planeta_s3 == -1{
		reloj_vector_s3 = {0,0,0}
	}else{
		reloj_vector_s3 = hash_s1[indice_planeta_s3].vector
	}

	conn.Close()

	ip = EscogerIP_MonotonicReads(reloj_vector_s1, reloj_vector_s2, reloj_vector_s3, reloj_vector_Leia)

	// ESCOGER IP QUE CUMPLA CON MONOTONIC READS

	// REALIZAR CONEXION CON SERVIDORES
	//ip = Server2Address

	conn2, err2 := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())

	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}

	// Client Stub to perform RPCs
	client2 := pb.NewComunicationClient(conn2)
	// Contact the server and psirint out its response.
	ctx2 := context.Background()

	r, _ := client2.Comands_Broker_Fulcrum(ctx2, &pb.ComandBFRequest{Operacion: operacion, NombrePlaneta: planeta, NombreCiudad: ciudad, Ip: ip})

	cant_soldados = r.GetCantRebelds()
	reloj_vector_Leia = r.GetRelojVector()

	conn2.Close()

	return &pb.ComandLBReply{CantRebelds: cant_soldados, RelojVector: reloj_vector_Leia}, nil
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
