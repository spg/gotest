package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	theMap := make(map[string]int)

	for _, word := range fields {
		theMap[word] = theMap[word] + 1
	}

	return theMap
}

func main() {
	wc.Test(WordCount)
}
