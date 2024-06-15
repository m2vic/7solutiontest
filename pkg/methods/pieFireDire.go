package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Meat(text string) []byte {

	// Clean the string and split into words
	text = strings.ToLower(text)
	cleanedText := strings.ReplaceAll(text, ".", "")
	cleanedText = strings.ReplaceAll(cleanedText, ",", "")
	words := strings.Fields(cleanedText)

	// Create a map to hold the count of each word
	wordCount := make(map[string]int)

	// Iterate over the words and count each occurrence
	for _, word := range words {
		wordCount[word]++
	}

	// Print the word count
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
	jsonData, err := json.Marshal(wordCount)
	if err != nil {
		return []byte(err.Error())
	}

	return jsonData

}
