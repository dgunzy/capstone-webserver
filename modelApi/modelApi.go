package modelApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ModelResponse struct {
	GeneratedText string `json:"generated_text"`
}

func SendModelRequest(message string) string {

	payload := map[string]string{
		"inputs": "summarize the key points clearly and concisely:  " + message,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return err.Error()
	}

	req, err := http.NewRequest("POST", "https://b1rd2yqmtubot0gq.us-east-1.aws.endpoints.huggingface.cloud", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error()
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer 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")

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

	fmt.Printf("Response: %s\n", string(responseBody))

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
	// var responseText ModelResponse

	// errrrr := json.Unmarshal(responseBody, &responseText)
	// if errrrr != nil {
	// 	fmt.Println("Error unmarshaling response:", err)

	// }

	// return string(responseText.GeneratedText)
}
