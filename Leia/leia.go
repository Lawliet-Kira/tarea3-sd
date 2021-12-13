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

		r, _ := client.Comands_Leia_Broker(ctx, &pb.ComandLBRequest{Operacion: operacion, NombrePlaneta: nombre_planeta, NombreCiudad: nombre_ciudad, RelojVector: reloj_vector_Leia})

		cant_soldados := r.GetCantRebelds()

		reloj_vector_Leia := r.GetRelojVector()

		fmt.Println("Cantidad de soldados: ", cant_soldados)
		fmt.Println("Reloj vector Leia: ", reloj_vector_Leia)

	}

}
