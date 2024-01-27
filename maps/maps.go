package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)

	stringArray := strings.Fields(s)

	fmt.Println(stringArray)

	for _, word := range stringArray {
		words[word]++
	}

	return words
}

func main() {
	wc.Test(WordCount)
}
