package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	var port string
	if len(os.Args) < 2 {
		port = "8000"
		log.Println("Using default port 8000")
	} else {
		port = os.Args[1]
		log.Println("Using port " + port)
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle multiple connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
