package modelApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ModelRequest struct {
	Inputs   string `json:"inputs"`
	NumBeams int    `json:"num_beams,omitempty"`
}

type ModelResponse struct {
	GeneratedText string `json:"summary_text"`
}

func SendModelRequest(message string, numBeams ...int) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(" no .env selected")
	}

	uri := os.Getenv("URI")
	auth := os.Getenv("AUTH")

	var nb int
	if len(numBeams) > 0 {
		nb = numBeams[0]
	}

	payload := ModelRequest{
		Inputs:   message,
		NumBeams: nb,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return err.Error()
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error()
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", auth)

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request", err)
		return err.Error()
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err.Error()
	}
	if response.StatusCode != http.StatusOK {
		return "Service Unavailable"
	}

	// fmt.Printf("Response: %s\n", string(responseBody))

	var responses []ModelResponse

	err = json.Unmarshal(responseBody, &responses)

	if err != nil {
		return err.Error()
	}
	if len(responses) > 0 {
		return responses[0].GeneratedText
	} else {
		return "No response received"
	}
}
