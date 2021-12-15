package main

import (
	"bufio"
	"context"
	"fmt"
	pb "lab3/game/helloworld"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"google.golang.org/grpc"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

const (
	BrokerAddress = "10.6.43.116:50051"
	defaultName   = "world"
)

var opcion = ""

type Keyvalue struct {
	planeta string
	vector  []int32
}

var relojes_Leia []Keyvalue
var ultimo_ip_server string
var registro_ciudades []string

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

func main() {
	// Crear un gRPC canal para comunicarse con el servidor
	// 	-> Esto se crea pasando server address y port number a grpc.Dial()
	conn, err := grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewComunicationClient(conn)
	ctx := context.Background()

	for opcion != "exit" {

		var operacion string
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

		comand = strings.TrimSuffix(comand, "\n")
		fmt.Println("after suffix: ", comand)

		splited := strings.Split(comand, " ")

		nombre_planeta := splited[0]
		nombre_ciudad := splited[1]

		fmt.Println("nombrePlaneta: ", nombre_planeta)
		fmt.Println("nombreCiudad: ", nombre_ciudad)

		indice_planeta := findHashing(relojes_Leia, nombre_planeta)

		if indice_planeta == -1 {
			Hashing = append(Hashing, {planeta: nombre_planeta, vector: int32{0,0,0}})
		}

		reloj_vector_L := Hashing[indice_planeta].vector

		r, _ := client.Comands_Leia_Broker(ctx, &pb.ComandLBRequest{Operacion: operacion, NombrePlaneta: nombre_planeta, NombreCiudad: nombre_ciudad, RelojVector: reloj_vector_L})

		cant_soldados := r.GetCantRebelds()

		reloj_vector_Leia := r.GetRelojVector()

		Hashing[indice_planeta].vector = reloj_vector_Leia

		registro_ciudades = append(registro_ciudades, nombre_ciudad)

		fmt.Println("Cantidad de soldados: ", cant_soldados)
		fmt.Println("Reloj vector Leia: ", reloj_vector_Leia)

	}

}
