package main

import (
	"os"

	"fmt"

	"golang.org/x/net/html"
)

func main() {
	var iBuffer [10]int
	slice := iBuffer[0:0]
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	//func add1(r rune) rune { return r + 1}

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
