package main

import (
	"fmt"
)

func normalizeURL(inputURL string) (string, error) {

	formattedString := ""

	for i, b := range inputURL {
		currentChar := string(b)
		if currentChar == ":" {
			protocol := inputURL[0:i]

			if protocol != "http" && protocol != "https" {
				return inputURL, fmt.Errorf("error with protocol")
			}

			formattedString = inputURL[i+3 : len(inputURL)]
		}
		if currentChar == "w" && string(inputURL[i+1]) == "w" && string(inputURL[i+2]) == "w" && string(inputURL[i+3]) == "." {
			formattedString = inputURL[i+4 : len(inputURL)]
		}
	}

	lastChar := string(inputURL[len(inputURL)-1])

	if lastChar == "/" {
		formattedString = formattedString[:len(formattedString)-1]
	}

	return formattedString, nil
}
