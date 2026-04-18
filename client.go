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
	//NewScanner(dial) attaches to the connection and watches the entire stream. 
	for read.Scan() {
		//Then, each time Scan() is called, it grabs the next individual line from that stream
		//so NewScanner covers the whole connection, while Scan() picks off one line at a time.
		fmt.Println(read.Text())
		 //listen.Text() reads and returns the data that Scan() just captured from the stream
	}
}
