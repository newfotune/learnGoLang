package main

/**
Basic Loops
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Main Called")
	findDup()
	fmt.Println("find dup returned")

	/*var s, sep string
	//basic loop through passed in argument
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
	s, sep = "", ""
	//range returns index and element at said index
	// use _ to ignore the index...as it is not needed
	for index, arg := range os.Args[0:] {
		fmt.Printf("%d - %s\n", index, arg)
	}*/

	//fmt.Println(s)
	//fmt.Println(strings.Join(os.Args[1:], "\n"))
}
func findDup() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println("Error encountered")
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	//ignoring potential errors from inout.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//Ignoring potential Errors from input.Err()
}
