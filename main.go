package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(connecting net.Conn) {
	defer connecting.Close()

	message := "hello world \n"
	connecting.Write([]byte(message))

}

func main() {
	listening, error := net.Listen("tcp", "0.0.0.0:8080")
	if error != nil {
		fmt.Println("error listenining")
		log.Fatal(error)
	}
	defer listening.Close()
	fmt.Println("server connected on 8080")

	for {
		connecting, error := listening.Accept()
		if error != nil {
			fmt.Println("error accepting")
			log.Fatal(error)
		}
		go handleConnection(connecting)
	}
}
