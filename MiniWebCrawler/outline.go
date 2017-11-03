package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	var path []string
	printOutline(doc, path)
	//m := make(map[string]int)
	//countElements(doc, &m)
	//fmt.Println(m)
}

func printOutline(node *html.Node, path []string) {

	if node.Type == html.ElementNode {
		path = append(path, strings.TrimSpace(node.Data))
		fmt.Println(path)
	}
	for node = node.FirstChild; node != nil; node = node.NextSibling {
		printOutline(node, path)
	}
}

func countElements(node *html.Node, m *map[string]int) {
	var themap = *m
	if node.Type == html.ElementNode || node.Type == html.DocumentNode {
		themap[node.Data]++
	}
	for node = node.FirstChild; node != nil; node = node.NextSibling {
		countElements(node, &themap)
	}
}
