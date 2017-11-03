package main

import (
	"fmt"
	"log"
	"os"
)

const file = "test.txt"
const aud = "01_Inception.mp3"

func main() {
	f, err := os.Open(aud)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer f.Close()
	buf := make([]byte, 5)

	for {
		_, err = f.Read(buf)
		if err != nil {
			//log.Fatalf("Error: %v\n", err)
			break
		}

		fmt.Printf("%v \n", buf)
	}
}
