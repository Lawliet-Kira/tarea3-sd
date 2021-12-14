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
var idFulcrum int

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
func findHashing(Hashing []Keyvalue, planeta string) int {

	for i, keyvalue := range Hashing {
		if keyvalue.planeta == planeta {
			fmt.Println("Planeta encontrado")
			return i
		}
	}

	return -1

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
		index := findHashing(Hashing, planeta)

		log.Println(Hashing[index].vector)

	} else if errors.Is(err, os.ErrNotExist) {

		// path/to/whatever does *not* exist
		createFile(path)
		Hashing = append(Hashing, *newKeyvalue(planeta))
		index := findHashing(Hashing, planeta)
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

	index := findHashing(Hashing, planeta)

	if index != -1 {
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

	index := findHashing(Hashing, planeta)

	if index != -1 {

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

	index := findHashing(Hashing, planeta)

	if index != -1 {
		return "success"
	}

	return "error"

}

func (s *server) Comands_Informantes_Fulcrum(ctx context.Context, in *pb.ComandIFRequest) (*pb.ComandIFReply, error) {

	// CHANGE
	reloj_vector_s := []int32{1, 0, 1}

	operacion := in.GetOperacion()
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	valor := in.GetValor()

	fmt.Println("operacion: ", operacion)
	fmt.Println("nameplanet: ", planeta)
	fmt.Println("namecity: ", ciudad)
	fmt.Println("value: ", valor)

	switch operacion {

	// LOGICA DE ADDCITY
	case "1":
		fmt.Println(AddCity(planeta, ciudad, valor))
		index := findHashing(Hashing, planeta)
		Hashing[index].vector[idFulcrum]++
	// LOGICA DE UPDATE NAME
	case "2":
		fmt.Println(UpdateName(planeta, ciudad, valor))
		index := findHashing(Hashing, planeta)
		Hashing[index].vector[idFulcrum]++
	// LOGICA UPDATE NUMBER
	case "3":
		fmt.Println(UpdateNumber(planeta, ciudad, valor))
		index := findHashing(Hashing, planeta)
		Hashing[index].vector[idFulcrum]++
	// LOGICA DELETE CITY
	case "4":
		fmt.Println(DeleteCity(planeta, ciudad))
		index := findHashing(Hashing, planeta)
		Hashing[index].vector[idFulcrum]++
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

	// LOGICA OPERACION GET
	cant_rebeldes := GetNumberRebelds(planeta, ciudad)

	return &pb.ComandBFReply{CantRebelds: cant_rebeldes, RelojVector: reloj_vector_s}, nil
}

func (s *server) Comands_Request_Hashing(ctx context.Context, in *pb.PingMsg) (*pb.HashRepply, error) {

	pbHashing := []*pb.HashRepply_KeyValue{}

	for _, keyvalue := range Hashing {
		temp := pb.HashRepply_KeyValue{Planeta: keyvalue.planeta, RelojVector: keyvalue.vector}
		pbHashing = append(pbHashing, &temp)
	}

	return &pb.HashRepply{Hashing: pbHashing}, nil
}

func (s *server) Comands_Request_Files(ctx context.Context, in *pb.PingMsg) (*pb.ComandFFFiles, error) {

	path, err := os.Getwd()

	var reloj_vector []int32

	if err != nil {
		log.Println(err)
	}

	logpath := path + "/logs/" + in.GetSignal() + ".txt"

	filepath := path + "/planetas/" + in.GetSignal() + ".txt"

	index := findHashing(Hashing, in.GetSignal())

	if index == -1 {

		// Agregar registro planeta y crear log vacio
		// Reloj [0,0,0]
		Hashing = append(Hashing, *newKeyvalue(in.GetSignal()))
		createFile(filepath)
		createFile(logpath)
		reloj_vector = []int32{0, 0, 0}

	} else {
		reloj_vector = Hashing[index].vector
	}

	//Enviar logs y reloj del planeta

	//Abrir archivo
	input, err := ioutil.ReadFile(logpath)

	if err != nil {
		log.Fatalln(err)
	}

	// Matrix con las lineas del archivo
	lines := strings.Split(string(input), "\n")

	return &pb.ComandFFFiles{Text: lines, RelojVector: reloj_vector}, nil

}
func (s *server) Comands_Retrieve_Files(ctx context.Context, in *pb.ComandFFFiles) (*pb.PingMsg, error) {

	//Replicar cambios recibidos del nodo dominante en el registro planetario borrar logs y guardar reloj
	//Leer texto recibido, sobreescribir registro y borrar log
	// Planeta
	target := in.GetPlaneta()
	text := in.GetText()
	path, _ := os.Getwd()
	filepath := path + "/planetas/" + target + ".txt"

	// Replace Vector Dominante
	Hashing[findHashing(Hashing, target)].vector = in.GetRelojVector()

	// Replace Planet Register
	text_replace := strings.Join(text, "\n")

	err := ioutil.WriteFile(filepath, []byte(text_replace), 0644)

	if err != nil {
		log.Fatal(err)
	}

	// Delete Logs

	path_logs := path + "/logs/" + target

	e := os.Remove(path_logs + ".txt")

	if e != nil {
		log.Fatal(e)
	}

	return &pb.PingMsg{Signal: ""}, nil

}

func ApplyChanges(pos int32, val int32, valDom int32, logs []string, target string) {

	if val > valDom {

		for _, accion := range logs {

			accion = strings.TrimSuffix(accion, "\n")
			split_line := strings.Split(accion, " ")
			operacion := split_line[0]
			planeta := split_line[1]
			ciudad := split_line[2]
			valor := ""

			if len(split_line) > 3 {
				valor = split_line[3]
			}

			switch operacion {
			case "AddCity":
				fmt.Println(AddCity(planeta, ciudad, valor))
			// LOGICA DE UPDATE NAME
			case "UpdateName":
				fmt.Println(UpdateName(planeta, ciudad, valor))
			// LOGICA UPDATE NUMBER
			case "UpdateNumber":
				fmt.Println(UpdateNumber(planeta, ciudad, valor))
			// LOGICA DELETE CITY
			case "DeleteCity":
				fmt.Println(DeleteCity(planeta, ciudad))
			}

			Hashing[findHashing(Hashing, target)].vector[pos]++

		}

	}

}

func ConsistenciaEventual() {

	fmt.Println("Consistencia Eventual...1")
	// Establecer conexión con servidor 3
	path, _ := os.Getwd()

	conn, err := grpc.Dial(Server3Address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewComunicationClient(conn)
	ctx := context.Background()

	defer conn.Close()

	// Establecer conexión con servidor 1
	conn2, err2 := grpc.Dial(Server1Address, grpc.WithInsecure(), grpc.WithBlock())

	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}

	client2 := pb.NewComunicationClient(conn2)
	ctx2 := context.Background()

	defer conn2.Close()

	for true {

		time.Sleep(60 * time.Second)
		fmt.Println("Consistencia Eventual...")

		//HACER PING
		signal := "Pingeao"
		fmt.Println("Pingeao")

		// Avisar a los servidores que envien sus Hashing
		r1, _ := client.Comands_Request_Hashing(ctx, &pb.PingMsg{Signal: signal})
		newHash := MergeHashing(Hashing, r1.GetHashing())

		r2, _ := client2.Comands_Request_Hashing(ctx2, &pb.PingMsg{Signal: signal})
		newHash = MergeHashing(newHash, r2.GetHashing())

		for _, keyvalue := range newHash {

			target := keyvalue.planeta
			filepath := path + "/planetas/" + target + ".txt"

			if findHashing(Hashing, target) == -1 {
				logpath := path + "/logs/" + target + ".txt"
				createFile(logpath)
				createFile(filepath)
				Hashing = append(Hashing, *newKeyvalue(target))
			}

			// Recuperación de Logs y Reloj para un Planeta Particular del SV esclavo S3
			r1, _ := client.Comands_Request_Files(ctx, &pb.PingMsg{Signal: target})
			fmt.Println("Logs S3: ", r1.GetText())
			fmt.Println("Reloj S3: ", r1.GetRelojVector())
			relojDom := Hashing[findHashing(Hashing, target)].vector // Reloj Dominante S2
			logs1 := r1.GetText()                                    // Logs del Esclavo S3
			reloj1 := r1.GetRelojVector()                            // Reloj del Esclavo S3

			// Recuperación de Logs y Reloj para un Planeta Particular del SV esclavo S1
			r2, _ := client2.Comands_Request_Files(ctx2, &pb.PingMsg{Signal: target})
			fmt.Println("Logs S1: ", r2.GetText())
			fmt.Println("Reloj S1: ", r2.GetRelojVector())
			logs2 := r2.GetText()         // Logs del Esclavo S1
			reloj2 := r2.GetRelojVector() // Reloj del Esclavo S1

			// Aplicar cambios del Log al registro planetario
			ApplyChanges(0, reloj1[0], relojDom[0], logs1, target) // sd: [0,0,0] se: [2,0,0]
			ApplyChanges(2, reloj2[2], relojDom[2], logs2, target) // sd: [0,0,0] se: [0,0,2]

			// Lectura del archivo de planeta
			//Abrir archivo
			input, err := ioutil.ReadFile(filepath)

			if err != nil {
				log.Fatalln(err)
			}

			// Matrix con las lineas del archivo
			linesReg := strings.Split(string(input), "\n")
			relojDom = Hashing[findHashing(Hashing, target)].vector

			// Replicación de registro de planeta y reloj sobre Servidores Esclavos
			client.Comands_Retrieve_Files(ctx, &pb.ComandFFFiles{Text: linesReg, RelojVector: relojDom, Planeta: target})

			client2.Comands_Retrieve_Files(ctx2, &pb.ComandFFFiles{Text: linesReg, RelojVector: relojDom, Planeta: target})

		}

	}

}

const (
	port           = ":50052"
	Server1Address = "10.6.43.113:50052"
	Server2Address = "10.6.43.114:50052"
	Server3Address = "10.6.43.115:50052"
)

///////////////////////////////////////////////////////////////////////////////////////////
//////////////////////// 	UTILIDADES 	 //////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////

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

//Fusiona la tabla de hashing de
func MergeHashing(Hash1 []Keyvalue, Hash2 []*pb.HashRepply_KeyValue) []Keyvalue {

	for _, keyvalue := range Hash2 {
		temp := Keyvalue{planeta: keyvalue.GetPlaneta(), vector: []int32{0, 0, 0}}
		Hash1 = append(Hash1, temp)
	}
	check := make(map[string]int)
	res := make([]Keyvalue, 0)
	for _, val := range Hash1 {
		check[val.planeta] = 1
	}

	for planeta, _ := range check {
		index := findHashing(Hash1, planeta)
		res = append(res, Hash1[index])
	}
	fmt.Println(res)
	return res
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
	switch ip {
	case "10.6.43.113":
		idFulcrum = 0
	case "10.6.43.114":
		idFulcrum = 1
	case "10.6.43.115":
		idFulcrum = 2
	}
	//Fulcrum dominante
	if idFulcrum == 1 {
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
