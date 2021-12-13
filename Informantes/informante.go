package main

import (
	"context"
	"fmt"
	pb "lab3/game/helloworld"
	"log"
	"os"
	"os/exec"
	"runtime"

	"google.golang.org/grpc"
)

const (
	BrokerAddress = "10.6.43.116:50051"
	defaultName   = "world"
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

var opcion = ""
var reloj_vector_Informante []int32

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
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println("Escoge a:")

	for opcion != "exit" {

		var operacion string
		var valor string
		var comand string

		// MENÚ
		fmt.Println("Ingrese Operación: ")
		fmt.Println("	1. Añadir Ciudad")
		fmt.Println("	2. Actualizar Nombre")
		fmt.Println("	3. Actualiar Numero")
		fmt.Println("	4. Eliminar Ciudad")

		fmt.Scanf("%s\n", &operacion)

		// clean the prompt
		CallClear()

		fmt.Println("\nIngresar comando: ")

		fmt.Scanf("%s\n", &comand)
		splited := comand.split(" ")

		fmt.Println("command: ", comand)
		fmt.Println("c")

		nombre_planeta := splited[0]
		nombre_ciudad := splited[1]

		fmt.Println("nombrePlaneta: ", nombre_planeta)
		fmt.Println("nombreCiudad: ", nombre_ciudad)

		if len(splited) == 3 {
			valor = splited[2]
		} else {
			valor = ""
		}

		r, _ := client.Comands_Informantes_Broker(ctx, &pb.ComandIBRequest{Operacion: operacion, NombrePlaneta: nombre_planeta, NombreCiudad: nombre_ciudad, Valor: valor, RelojVector: reloj_vector_Informante})

		fmt.Println("Direccion IP seleccionada: ", r)

		// Connection with IP Fulcrum

		FulcrumAddress := r.GetIp()

		conn2, err2 := grpc.Dial(FulcrumAddress, grpc.WithInsecure(), grpc.WithBlock())

		if err2 != nil {
			log.Fatalf("did not connect: %v", err2)
		}

		client2 := pb.NewComunicationClient(conn2)

		fmt.Println("Escribe el comando: ")

		r2, _ := client2.Comands_Informantes_Fulcrum(ctx, &pb.ComandIFRequest{Operacion: operacion, NombrePlaneta: nombre_planeta, NombreCiudad: nombre_ciudad, Valor: valor})

		fmt.Println("Reply: ", r2)

		// Close connection with Fulcrum
		conn2.Close()

	}

}
