package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)


func main() {
	var elements = make(map[string]int)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count: %v\n", err)
		os.Exit(1)
	}
	countElements(elements, doc)
	for element, count := range elements {
		fmt.Printf("%s: %d\n", element, count)
	}
}

func countElements(elements map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	if c := n.FirstChild; c != nil {
		countElements(elements, c)
	}

	if s := n.NextSibling; s != nil {
		countElements(elements, s)
	}
}
