package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	grpc "./grpc"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func enviar(records [][]string, tipo string, segundos int) {
	var aux string
	for i, producto := range records {
		if i != 0 {
			for j, carac := range producto {
				if j == 0 {
					aux = tipo + "+" + carac + "+"
				} else if j != 4 && tipo == "retail" {
					aux = aux + carac + "+"
				} else if j == 4 && tipo == "retail" {
					aux = aux + carac
				} else if j != 5 && tipo == "pyme" {
					aux = aux + carac + "+"
				} else if j == 5 && tipo == "pyme" {
					aux = aux + carac
				}
			}
			time.Sleep(time.Duration(segundos) * time.Second)
			grpc.Send(aux)
		}
	}
}

func main() {
	fmt.Println("Bienvenido al sistema de Clientes!\n Seleccione una opción:")

	var opcion string
	opcion = "0"

	for opcion != "4" {

		fmt.Println("	1) Ingresar órdenes como Cliente retail\n	2) Ingresar órdenes como Cliente pyme\n	3) Ver estado de un pedido\n	4) Salir")
		fmt.Scanf("%s", &opcion)
		var segundos int
		if opcion == "1" {
			fmt.Println("Elegiste la opción ingresar ordenes como cliente retail")
			fmt.Println("Indique cada cuántos segundos quiere enviar órdenes:")
			fmt.Scanf("%d", &segundos)
			records := readCsvFile("archivos/retail.csv")
			enviar(records, "retail", segundos)
		} else if opcion == "2" {
			fmt.Println("Elegiste la opción ingresar ordenes como cliente pyme")
			fmt.Println("Indique cada cuántos segundos quiere enviar órdenes:")
			fmt.Scanf("%d", &segundos)
			records := readCsvFile("archivos/pymes.csv")
			enviar(records, "pyme", segundos)
		} else if opcion == "3" {
			fmt.Println("Elegiste la opción ver estado de una orden")
		} else if opcion == "4" {
			fmt.Print("Adios\n")
		} else {
			fmt.Println("Por favor, selecciona una opción válida")
		}
	}
}
