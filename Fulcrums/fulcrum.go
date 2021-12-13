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
	"path/filepath"
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
func findHashing(planeta string) int32 {

	var cont int32 = 0

	for _, planet := range Hashing {
		if planet.planeta == planeta {
			return cont
		}
		cont += 1
	}
	return -1

}

// Verifica si existe un archivo, en caso contrario lo crea
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

// Temporary directory
const tmpDir = "planetas/"

func AddCity(planeta string, ciudad string, valor string) string {

	// "nombre_planeta nombre_ciudad [nuevo_valor]"

	path := filepath.Join(tmpDir, planeta+".txt")

	// VERIFICAR QUE EL ARCHIVO EXISTE
	if _, err := os.Stat(path); err == nil {

		// path/to/whatever exists
		index := findHashing(planeta)
		Hashing[index].vector[idFulcrum]++

	} else if errors.Is(err, os.ErrNotExist) {

		// path/to/whatever does *not* exist
		createFile(path)
		Hashing = append(Hashing, *newKeyvalue(planeta))
		index := findHashing(planeta)
		Hashing[index].vector[idFulcrum]++

	}

	// Abrir archivo y escribir al final
	if valor == "" {
		valor = "0"
	}

	dataByte := []byte(planeta + " " + ciudad + " " + valor + "\n")

	// Escribir contenido en el archivo
	err := ioutil.WriteFile(path, dataByte, 0777)

	// log Error
	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	result := "success"

	return result

}

func UpdateName(planeta string, ciudad string, valor string) string {

	path := filepath.Join(tmpDir, "/"+planeta+".txt")

	//Abrir archivo
	input, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	// Matrix con las lineas del archivo
	lines := strings.Split(string(input), "\n")

	// Leer linea x linea
	for i, line := range lines {
		if strings.Contains(line, ciudad) {
			split_line := strings.Split(line, " ")
			np := split_line[0]
			v := split_line[2]
			lines[i] = np + " " + valor + " " + v
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(path, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	index := findHashing(planeta)
	Hashing[index].vector[idFulcrum]++

	return "success"

}

func UpdateNumber(planeta string, ciudad string, valor string) string {

	path := filepath.Join(tmpDir, "/"+planeta+".txt")

	// Abrir archivo
	input, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
		return "error"
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

	err = ioutil.WriteFile(path, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	index := findHashing(planeta)
	Hashing[index].vector[idFulcrum]++

	return "success"

}

func DeleteCity(planeta string, ciudad string) string {

	path := filepath.Join(tmpDir, "/"+planeta+".txt")

	//Abrir archivo
	input, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	lines := strings.Split(string(input), "\n")

	var del_index int

	for i, line := range lines {
		if strings.Contains(line, ciudad) {
			del_index = i
		}
	}

	lines_to_write := append(lines[:del_index], lines[del_index+1:]...)

	output := strings.Join(lines_to_write, "\n")

	err = ioutil.WriteFile(path, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	index := findHashing(planeta)
	Hashing[index].vector[idFulcrum]++

	return "success"

}

func (s *server) Comands_Informantes_Fulcrum(ctx context.Context, in *pb.ComandIFRequest) (*pb.ComandIFReply, error) {

	// CHANGE
	reloj_vector_s := []int32{1, 0, 1}

	operacion := in.GetOperacion()
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	valor := in.GetValor()
	localip := in.GetIp()

	switch localip {
	case "10.6.43.113:50052":
		idFulcrum = 0
	case "10.6.43.114:50052":
		idFulcrum = 1
	case "10.6.43.115:50052":
		idFulcrum = 2
	}

	fmt.Println("operacion: ", operacion)
	fmt.Println("nameplanet: ", planeta)
	fmt.Println("namecity: ", ciudad)
	fmt.Println("value: ", valor)

	switch operacion {

	// LOGICA DE ADDCITY
	case "1":
		AddCity(planeta, ciudad, valor)

	// LOGICA DE UPDATE NAME
	case "2":
		UpdateName(planeta, ciudad, valor)

	// LOGICA UPDATE NUMBER
	case "3":
		UpdateNumber(planeta, ciudad, valor)

	// LOGICA DELETE CITY
	case "4":
		DeleteCity(planeta, ciudad)

	}

	// VERIFICAR SI EXISTE EL ARCHIVO LOGS DE REGISTROS

	// DEVOLVER RELOJ ARCHIVO
	index := findHashing(planeta)
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
