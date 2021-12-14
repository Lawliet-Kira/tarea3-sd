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
	"time"

	"google.golang.org/grpc"
)

var idFulcrum int

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
func findHashing(planeta string) int {

	for i, keyvalue := range Hashing {
		if keyvalue.planeta == planeta {
			fmt.Println("Planeta encontrado")
			return i
		}
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

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func AddCity(planeta string, ciudad string, valor string) string {

	// "nombre_planeta nombre_ciudad [nuevo_valor]"
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	path = path + "/planetas/" + planeta + ".txt"

	fmt.Println("current path: ", path)

	// VERIFICAR QUE EL ARCHIVO EXISTE
	if _, err := os.Stat(path); err == nil {

		// path/to/whatever exists
		fmt.Println("El archivo existe")
		index := findHashing(planeta)
		Hashing[index].vector[idFulcrum]++
		log.Println(Hashing[index].vector)

	} else if errors.Is(err, os.ErrNotExist) {

		// path/to/whatever does *not* exist
		createFile(path)
		Hashing = append(Hashing, *newKeyvalue(planeta))
		index := findHashing(planeta)
		Hashing[index].vector[idFulcrum]++
		log.Println(Hashing[index].vector)

	}

	// Abrir archivo y escribir al final
	if valor == "" {
		valor = "0"
	}

	text := planeta + " " + ciudad + " " + valor + "\n"

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		log.Fatalln(err)
		return "error"
	}

	return "success"

}

func UpdateName(planeta string, ciudad string, valor string) string {

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "error"
	}

	path = path + "/planetas/" + planeta + ".txt"

	fmt.Println("current path: ", path)

	//Abrir archivo
	input, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	// Matrix con las lineas del archivo
	lines := strings.Split(string(input), "\n")

	fmt.Println("Nueva_ciudad: ", valor, "-----")

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

	if index != -1 {
		Hashing[index].vector[idFulcrum]++
		log.Println(Hashing[index].vector)
		return "success"
	}

	return "error"

}

func UpdateNumber(planeta string, ciudad string, valor string) string {

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "error"
	}

	path = path + "/planetas/" + planeta + ".txt"

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

	if index != -1 {
		Hashing[index].vector[idFulcrum]++
		log.Println(Hashing[index].vector)
		return "success"
	}

	return "error"

}

func DeleteCity(planeta string, ciudad string) string {

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "error"
	}

	path = path + "/planetas/" + planeta + ".txt"

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

	if index != -1 {
		Hashing[index].vector[idFulcrum]++
		log.Println(Hashing[index].vector)
		return "success"
	}

	return "error"

}

func EscribirLog(operacion string, planeta string, ciudad string, valor string) string {

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "error"
	}

	path = path + "/logs/" + planeta + ".txt"

	// VERIFICAR SI EXISTE EL ARCHIVO, SINO CREARLO
	createFile(path)

	//ESCRIBIR EL LOG AL FINAL DEL ARCHIVO
	op := ""

	switch operacion {

	case "1":
		op = "AddCity"
	case "2":
		op = "UpdateName"
	case "3":
		op = "UpdateNumber"
	case "4":
		op = "DeleteCity"
	}

	if operacion == "1" && valor == "" {
		valor = "0"
	}

	ciudad = strings.TrimSuffix(ciudad, "\n")
	valor = strings.TrimSuffix(valor, "\n")

	text := op + " " + planeta + " " + ciudad + " " + valor + "\n"

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		log.Fatalln(err)
		return "error"
	}

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
		fmt.Println(AddCity(planeta, ciudad, valor))

	// LOGICA DE UPDATE NAME
	case "2":
		fmt.Println(UpdateName(planeta, ciudad, valor))

	// LOGICA UPDATE NUMBER
	case "3":
		fmt.Println(UpdateNumber(planeta, ciudad, valor))

	// LOGICA DELETE CITY
	case "4":
		fmt.Println(DeleteCity(planeta, ciudad))
	}

	EscribirLog(operacion, planeta, ciudad, valor)

	// VERIFICAR SI EXISTE EL ARCHIVO LOGS DE REGISTROS

	// DEVOLVER RELOJ ARCHIVO
	/*index := findHashing(planeta)
	reloj_vector_s = Hashing[index].vector*/

	return &pb.ComandIFReply{RelojVector: reloj_vector_s}, nil

}

func GetNumberRebelds(planeta string, ciudad string) string {

	var cant_rebeldes = ""

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "error"
	}

	path = path + "/planetas/" + planeta + ".txt"

	//Abrir archivo
	input, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
		return "error"
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		fmt.Println("line ->(", line, ")")
		fmt.Println("ciudad ->(", ciudad, ")")
		if strings.Contains(line, ciudad) {
			split_line := strings.Split(line, " ")
			fmt.Println("1 ->", split_line[0])
			fmt.Println("2 ->", split_line[1])
			fmt.Println("3 ->", split_line[2])
			cant_rebeldes = split_line[2]
			break
		}
	}

	return cant_rebeldes
}

func (s *server) Comands_Broker_Fulcrum(ctx context.Context, in *pb.ComandBFRequest) (*pb.ComandBFReply, error) {

	// CHANGE
	reloj_vector_s := []int32{1, 2, 1}
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	localip := in.GetIp()

	switch localip {
	case "10.6.43.113:50052":
		idFulcrum = 0
	case "10.6.43.114:50052":
		idFulcrum = 1
	case "10.6.43.115:50052":
		idFulcrum = 2
	}

	// LOGICA OPERACION GET
	cant_rebeldes := GetNumberRebelds(planeta, ciudad)

	return &pb.ComandBFReply{CantRebelds: cant_rebeldes, RelojVector: reloj_vector_s}, nil
}

/*func (s *server) Comands_Fulcrum_Fulcrum(ctx context.Context, in *pb.ComandFFRequest) (*pb.ComandFFReply, error) {

	return &pb.ComandFFReply{RelojVector: reloj_vector_s}, nil
}*/

func ConsistenciaEventual() {

	for true {
		time.Sleep(10 * time.Second)
		fmt.Println("Consistencia Eventual...")
		//MENSAJE CONSISTENCIA EVENTUAL

	}

}

const (
	port = ":50052"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

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

	var ip string = GetLocalIP()

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}

	fmt.Println("localip: (", ip, ")")
	//Fulcrum dominante
	if string(ip) == "10.6.43.114" {
		fmt.Println("Soy el Fulcrum dominante uwu")
		// Function each seconds
		go ConsistenciaEventual()
	}

	// Llamar Server() con los detalles de puerto para realizar un bloquero
	// de espera hasta que el proceso sea killed o Stop() es llamado
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
