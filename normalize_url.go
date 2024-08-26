package main

import "fmt"

func normalizeURL(inputURL string) string {

	formattedString := ""

	for i, b := range inputURL {
		currentChar := string(b)
		if currentChar == ":" {
			formattedString = inputURL[i+3 : len(inputURL)]
			fmt.Println("change string ", formattedString)
		}
		if currentChar == "w" && string(inputURL[i+1]) == "w" && string(inputURL[i+2]) == "w" && string(inputURL[i+3]) == "." {
			formattedString = inputURL[i+4 : len(inputURL)]
		}
	}

	lastChar := string(inputURL[len(inputURL)-1])

	if lastChar == "/" {
		formattedString = formattedString[:len(formattedString)-1]
	}

	return formattedString
}
