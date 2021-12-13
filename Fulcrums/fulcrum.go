package main

import (
	"context"
	pb "lab3/game/helloworld"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComunicationServer
}
type Location struct {
	planeta string
	vector  []int32
}

func createFile(path string) {

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {

		var file, err = os.Create(path)

		if err != nil {
			log.Fatalf("Error al crear archivo: %v", err)
			return
		}

		defer file.Close()

	}

}

func AddCity(result string) string {

	// "nombre_planeta nombre_ciudad [nuevo_valor]"

	// VERIFICAR QUE EL ARCHIVO EXISTE

	return result
}

func UpdateName(result string) string {
	return result
}

func UpdateNumber(result string) string {
	return result
}

func DeleteCity(result string) string {
	return result
}

func (s *server) Comands_Informantes_Fulcrum(ctx context.Context, in *pb.ComandIFRequest) (*pb.ComandIFReply, error) {

	reloj_vector_s := []int32{1, 0, 1}

	operacion := in.GetOperacion()
	nombre_planeta := in.GetNombrePlaneta()
	nombre_ciudad := in.GetNombreCiudad()
	valor := in.GetValor()

	switch operacion {

	// LOGICA DE ADDCITY
	case "1":
		AddCity(nombre_planeta, nombre_ciudad, valor)

	// LOGICA DE UPDATE NAME
	case "2":
		UpdateName(nombre_planeta, nombre_ciudad, valor)

	// LOGICA UPDATE NUMBER
	case "3":
		UpdateNumber(nombre_planeta, nombre_ciudad, valor)

	// LOGICA DELETE CITY
	case "4":
		DeleteCity(nombre_planeta, nombre_ciudad)

	}

	// VERIFICAR SI EXISTE EL ARCHIVO LOGS DE REGISTROS

	return &pb.ComandIFReply{RelojVector: reloj_vector_s}, nil

}

const (
	port = ":50052"
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

	log.Printf("Server Fulcrum escuchando en %v", lis.Addr())

	// Llamar Server() con los detalles de puerto para realizar un bloquero
	// de espera hasta que el proceso sea killed o Stop() es llamado
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
