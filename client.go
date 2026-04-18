package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	dial, error := net.Dial("tcp", "127.0.0.1:8080")
	if error != nil {
		fmt.Println("error dialing")
		log.Fatal(error)
	}

	read := bufio.NewScanner(dial)
	for read.Scan() {
		fmt.Println(read.Text())
	}
}
