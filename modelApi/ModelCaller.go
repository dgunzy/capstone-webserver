package modelApi

import (
	"sort"
	"strings"
	"sync"
)

type indexSummary struct {
	index   int
	summary string
}

func ModelCaller(input string, chunkSize int) string {
	response := SendModelRequest("Are you awake?")

	if response == "Service Unavailable" {
		return "Service Unavailable"
	}

	if len(input) < 1200 {
		return SendModelRequest("summarize all key points: " + input)
	}
	chunks := ChunkText(input, chunkSize)
	// This used to be a channel of strings, but they would be unsorted, so they are a channel of structs that contain an index to sort
	summaryChannel := make(chan indexSummary, len(chunks))
	var waitGroup sync.WaitGroup

	for i, chunk := range chunks {
		waitGroup.Add(1)

		go func(index int, chunk string) {
			defer waitGroup.Done()
			// fmt.Println("Chunk: \n" + chunk)
			summaryChannel <- indexSummary{index: index, summary: SendModelRequest("summarize: " + chunk)}
		}(i, chunk)
	}

	waitGroup.Wait()
	close(summaryChannel)

	var summarizedChunks []indexSummary
	for summary := range summaryChannel {
		summarizedChunks = append(summarizedChunks, summary)
	}
	sort.Slice(summarizedChunks, func(i, j int) bool {
		return summarizedChunks[i].index < summarizedChunks[j].index
	})

	var orderedSummaries []string
	for _, summary := range summarizedChunks {
		orderedSummaries = append(orderedSummaries, summary.summary)
	}

	combinedSummary := strings.Join(orderedSummaries, " ")

	// fmt.Println("Combine summary is: " + combinedSummary)
	if len(combinedSummary) < 1200 {
		return SendModelRequest("summarize all key points: " + combinedSummary)
	}
	return ModelCaller(combinedSummary, 1200)
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
