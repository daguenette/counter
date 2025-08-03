package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./words.txt")
	if err != nil {
		fmt.Println(err)
	}

	var wordCount int
	words := strings.SplitSeq(string(data), " ")
	for range words {
		wordCount++
	}

	fmt.Println(wordCount)
}
