package main

import (
	"fmt"
)

func Compress(input string) []int {
	// Initialize the dictionary with all possible single characters
	dict := make(map[string]int)
	for i := 0; i < 256; i++ {
		dict[string(i)] = i
	}

	compressedResult := []int{}
	buffer := ""

	for _, ch := range input {
		newBuffer := buffer + string(ch)
		// Check if the new string is in the dictionary
		if _, ok := dict[newBuffer]; ok {
			buffer = newBuffer
		} else {
			// Add the code for the previous buffer to the compressed result
			compressedResult = append(compressedResult, dict[buffer])

			// Add the new buffer to the dictionary
			dict[newBuffer] = len(dict)

			// Reset the buffer to the current character
			buffer = string(ch)
		}
	}

	// Add the final code for the buffer to the compressed result
	if buffer != "" {
		compressedResult = append(compressedResult, dict[buffer])
	}

	return compressedResult
}

func main() {
	input := "vgh"
	compressed := Compress(input)
	fmt.Println(compressed)
}
