package main

import (
	"fmt"
	"goLearn/src/tempconv"
	"os"
	"strconv"
	"time"
)

//Currency - type with underlying value int
type Currency int

const (
	//USD - United Stated Dollars
	USD Currency = iota
	//EUR - Euros
	EUR
	//GBP - Great Britain Pounds
	GBP
	//RMB - China's Ramini
	RMB
)

//Person - Struct that Represents a person in the Hospital
type Person struct {
	ID            int
	Name, Address string
	DoB           time.Time
}

var ceo Person

func main() {
	currentTime := time.Now()
	ceo = Person{ID: 123, Name: "Fortune", Address: "28012 33rd Avenue South", DoB: currentTime}
	ceo.printPerson()
	//symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	spring := months[4:7]
	fmt.Println(len(spring))
	for month := range spring {
		fmt.Printf("%s ", spring[month])
	}
	fmt.Println()
	s := []int{0, 1, 2, 3, 4, 5}
	//rotate s by n places to the right
	reverse(s)
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s[:2])
	fmt.Println(s)
	//doTempComv()
}

func (p Person) printPerson() {
	fmt.Printf("ID = %d, Name = %s, Address = %s, DoB = %v", p.ID, p.Name, p.Address, p.DoB)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func doTempComv() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	}
}
