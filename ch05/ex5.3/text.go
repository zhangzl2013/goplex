package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
	"strings"
)


func main() {
	var texts []string

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count: %v\n", err)
		os.Exit(1)
	}
	texts = countElements(texts, doc)

	for _, text := range texts {
		fmt.Printf("%s\n", text)
	}
}

func countElements(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode &&
		(n.Data == "script" || n.Data == "style" ){
		return texts
	}

	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		texts = append(texts, strings.TrimSpace(n.Data))
	}

	if c := n.FirstChild; c != nil {
		texts = countElements(texts, c)
	}

	if s := n.NextSibling; s != nil {
		texts = countElements(texts, s)
	}
	return texts
}
