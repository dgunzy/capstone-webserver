package modelApi

import (
	"fmt"
	"strings"
)

func ModelCaller(input string, chunkSize int) string {
	if len(input) < 1200 {
		return SendModelRequest(input)
	}
	chunks := ChunkText(input, chunkSize)

	var summerizedChunks []string
	for _, chunk := range chunks {
		fmt.Println("Chunk: \n" + chunk)
		summerizedChunks = append(summerizedChunks, SendModelRequest(chunk))
	}
	combinedSummary := strings.Join(summerizedChunks, " ")

	if len(combinedSummary) < 1200 {
		return SendModelRequest(combinedSummary)
	}
	ModelCaller(combinedSummary, 1200)
	return "Error chunking text in modelCaller"
}

func ChunkText(text string, chunkSize int) []string {
	words := strings.Fields(text)
	var chunks []string
	var currentChunk []string

	currentLen := 0
	for _, word := range words {
		if currentLen+len(word)+1 > chunkSize {
			chunks = append(chunks, strings.Join(currentChunk, " "))
			currentChunk = []string{word}
			currentLen = len(word)
		} else {
			currentChunk = append(currentChunk, word)
			currentLen += len(word) + 1
		}
	}
	chunks = append(chunks, strings.Join(currentChunk, " "))
	return chunks
}
