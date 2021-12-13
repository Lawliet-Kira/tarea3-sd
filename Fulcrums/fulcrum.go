package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	pb "lab3/game/helloworld"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComunicationServer
}

//Struct utilizado para guardar un planeta con su respectivo reloj vector
type Keyvalue struct {
	planeta string
	vector  []int32
}

//Lista de Keyvalues para guardar los relojes de cada planeta
var Hashing []Keyvalue

//Crea una nueva Keyvalue planeta, vector, inicializando este ultimo en [0,0,0]
func newKeyvalue(planeta string) *Keyvalue {
	var vector []int32
	vector = append(vector, 0)
	vector = append(vector, 0)
	vector = append(vector, 0)
	new := Keyvalue{planeta: planeta, vector: vector}
	return &new
}

//Encuentra la posicion en la que se encuentra el planeta deseado en la lista de Keyvalues
func findHashing(hashtable []Keyvalue, planeta string) []int32 {
	var index int
	var cont int32 = 0
	for _, planet := range hashtable {
		if planet.planeta == planeta {
			return cont
		}
		cont += 1
	}
	return -1
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

func AddCity(planeta string, ciudad string, valor string) resustring {

	// "nombre_planeta nombre_ciudad [nuevo_valor]"

	// VERIFICAR QUE EL ARCHIVO EXISTE
	if _, err := os.Stat("/planetas/" + planeta + ".txt"); err == nil {
		// path/to/whatever exists
		index := findHashing(Hashing, planeta)
		Hashing[index].vector[idFulcrum]++

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		createFile("/planetas/" + planeta + ".txt")

		Hashing = append(Hashing, *newKeyvalue(planeta))
		index := findHashing(Hashing, planeta)
		Hashing[index].vector[idFulcrum]++
	}

	//Abrir archivo y escribir al final
	if valor == nil {
		valor = 0
	}
	var f, err = os.OpenFile("/planetas/"+planeta+".txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "error"
	}

	n, err := f.WriteString(planeta + ciudad + valor + "\n")

	if err != nil {
		return "Errr "
	}

	result := "success"

	return result
}

func UpdateName(planeta string, ciudad string, valor string) string {

	//Abrir archivo
	input, err := ioutil.ReadFile("/planetas/" + planeta + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	//LEER LINEA POR LINEA

	for i, line := range lines {
		if strings.Contains(line, ciudad) {
			split_line := strings.Split(line, " ")
			np := split_line[0]
			v := split_line[2]
			lines[i] = np + " " + valor + " " + v
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile("/planetas/"+planeta+".txt", []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}

	index := findHashing(Hashing, planeta)
	Hashing[index].vector[idFulcrum]++
	result := "success"

	return result

}

func UpdateNumber(planeta string, ciudad string, valor string) string {

	// Abrir archivo
	input, err := ioutil.ReadFile("/planetas/" + planeta + ".txt")

	if err != nil {
		log.Fatalln(err)
	}

	// Arreglo de lineas
	lines := strings.Split(string(input), "\n")

	//LEER LINEA POR LINEA

	for i, line := range lines {
		if strings.Contains(line, ciudad) {
			split_line := strings.Split(line, " ")
			np := split_line[0]
			nc := split_line[1]
			lines[i] = np + " " + nc + " " + valor
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile("/planetas/"+planeta+".txt", []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}

	index := findHashing(Hashing, planeta)
	Hashing[index].vector[idFulcrum]++

	result := "success"

	return result

}

func DeleteCity(planeta string, ciudad string) string {

	//Abrir archivo
	input, err := ioutil.ReadFile("/planetas/" + planeta + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	var line_to_delete int32

	for i, line := range lines {
		if strings.Contains(line, ciudad) {
			line_to_delete = i
		}
	}

	lines_to_write := append(lines[:line_to_delete], lines[line_to_delete+1:]...)

	output := strings.Join(lines_to_write, "\n")

	err = ioutil.WriteFile("/planetas/"+planeta+".txt", []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}

	index := findHashing(Hashing, planeta)
	Hashing[index].vector[idFulcrum]++

	result := "success"

	return result
}

func (s *server) Comands_Informantes_Fulcrum(ctx context.Context, in *pb.ComandIFRequest) (*pb.ComandIFReply, error) {

	reloj_vector_s := []int32{1, 0, 1}

	operacion := in.GetOperacion()
	nombre_planeta := in.GetNombrePlaneta()
	nombre_ciudad := in.GetNombreCiudad()
	valor := in.GetValor()
	localip := in.GetIp()

	if idFulcrum == nil {
		switch localip {
		case "10.6.43.113:50052":
			idFulcrum = 0
		case "10.6.43.114:50052":
			idFulcrum = 1
		case "10.6.43.115:50052":
			idFulcrum = 2
		}
	}
	fmt.Println("operacion: ", operacion)
	fmt.Println("nameplanet: ", nombre_planeta)
	fmt.Println("namecity: ", nombre_ciudad)
	fmt.Println("value: ", valor)

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

	// DEVOLVER RELOJ ARCHIVO

	index := findHashing(Hashing, planeta)
	reloj_vector_s = Hashing[index].vector

	return &pb.ComandIFReply{RelojVector: reloj_vector_s}, nil

}

const (
	port = ":50052"
)

var idFulcrum int

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
