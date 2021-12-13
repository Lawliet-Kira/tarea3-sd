const (
	address     = "10.6.43.113:50051"
	defaultName = "world"
)

func main() {
	// Crear un gRPC canal para comunicarse con el servidor
	// 	-> Esto se crea pasando server address y port number a grpc.Dial()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Client Stub to perform RPCs
	client := pb.NewGameClient(conn)
	message := "HOLA deseo unirme a the game"
	// Contact the server and psirint out its response.
	name := "Jugador 1"
	var id int32 = 0
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
}