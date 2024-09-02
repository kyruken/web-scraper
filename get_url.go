package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	reader := strings.NewReader(htmlBody)

	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return urls, err
	}

	readNode(doc, "")

	fmt.Println("HTML parsed successfully!")

	return urls, nil
}

func readNode(n *html.Node, indent string) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		readNode(child, "")
	}
}
