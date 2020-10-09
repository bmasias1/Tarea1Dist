package grpc

import (
	"context"
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func Send(aEnviar string) {
	var conn *grpc.ClientConn
	//fmt.Println("Ingresa la ip")
	var addr string
	//fmt.Scanf("%s", &addr)
	addr = "10.10.28.154:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: aEnviar,
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error al llama a SayHello: %s", err)
	}

	log.Printf("Respuesta del Servidor: %s", response.Body)
}
