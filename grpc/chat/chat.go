package chat

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"encoding/csv"
	"io/ioutil"
	"strings"

	"golang.org/x/net/context"
)

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func EscribirCSV(aEscribir string) string {

	data, err := ioutil.ReadFile("../../archivos/indexAct.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))

	f, err := os.OpenFile("../../archivos/results.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return "0"
	}
	w := csv.NewWriter(f)
	dataStr := string(data[:])
	aEscribir = dataStr + "+" + aEscribir
	nueva := strings.Split(aEscribir, "+")
	retorno := "El código de seguimiento de " + nueva[2] + " es " + nueva[0]
	w.Write(nueva)

	dataInt, err := strconv.Atoi(dataStr)
	bs := []byte(strconv.Itoa(dataInt + 1))

	error := ioutil.WriteFile("../../archivos/indexAct.data", bs, 0777)
	if error != nil {
		fmt.Println(err)
	}

	w.Flush()
	return retorno
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje recibido desde el cliente: %s", message.Body)
	mensaje := EscribirCSV(message.Body)
	return &Message{Body: mensaje}, nil
}
