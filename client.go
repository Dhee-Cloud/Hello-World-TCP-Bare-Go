package main 

import (
	"net"
	"fmt"
	"io"
	"log"
)

func main () {
dial, error := net.Dial("tcp", "192.168.20.48:8080")
if error != nil {
	fmt.Println("error dialing")
	log.Fatal(error)
}	

read, error := io.ReadAll(dial)
if error != nil {
	fmt.Println("error reading")
	log.Fatal(error)
}
fmt.Println(string(read))
}