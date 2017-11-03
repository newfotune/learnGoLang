package main

import (
	"fmt"
)

type ByteCounter int

func main() {
	var c ByteCounter
	fmt.Println(c)
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	//convert int to byteCounter
	*c += ByteCounter(len(p))

	return len(p), nil
}
