package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	var buf []byte
	var reader func() bool

	reader = func() bool {
		bytesRead, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			log.Println(err)
			return false
		}
		log.Println("Read ", bytesRead, " bytes")
		return true
	}

	fmt.Printf("%v\b", reader())

	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		_, err = io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		if err != nil {
			log.Printf("err = %v\n", err)
		}
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
