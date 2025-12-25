package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Runner struct {
	ApiKey        string
	Client        *http.Client
	LLMParameters *Parameters
	Config        Config
}

type Parameters struct {
	LLMMessages         []Message `json:"messages"`
	Model               string    `json:"model"`                 // Selects the LLM running on Groq’s LPU infrastructure.
	Temperature         float32   `json:"temperature"`           // Controls randomness (Lower = more deterministic, Higher = more creative).
	MaxCompletionTokens uint      `json:"max_completion_tokens"` // Maximum tokens the model can generate in the response.
	TopP                float32   `json:"top_p"`                 // The model chooses from tokens whose cumulative probability ≤ top_p (1.0 = use all).
	Stream              bool      `json:"stream"`                // Whether to stream output or not.
	ReasoningEffort     string    `json:"reasoning_effort"`      // Controls how much internal reasoning depth the model applies. (low, medium, high)
	Stop                string    `json:"stop"`                  // As soon as the model produces this string in a row, generation ends.
}

type Message struct {
	Role    string `json:"role"`    // Instructions or rules for the model.
	Content string `json:"content"` // Your actual prompt.
}

func DefaultLLMParameters() *Parameters {
	return &Parameters{
		LLMMessages:         []Message{},
		Model:               "",
		Temperature:         0.7,      // Balanced creativity
		MaxCompletionTokens: 1024,     // Reasonable default
		TopP:                0.9,      // Standard value
		Stream:              false,    // Non-streaming by default
		ReasoningEffort:     "medium", // Balanced reasoning
		Stop:                "",       // No stop sequence by default
	}
}

func (r *Runner) RunGrok() ([]byte, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate runner:%v", err)
	}
	postData, err := json.Marshal(r.LLMParameters)
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal LLM Parameters to json: %v", err)
	}

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+r.ApiKey)
	if r.LLMParameters.Stream {
		headers.Set("Accept", "text/event-stream")
	}

	request := Request{
		r.Config.ApiBaseUrl.Groq,
		headers,
		r.Client,
	}

	resp, err := request.SendPost(bytes.NewReader(postData))
	if err != nil {
		return nil, fmt.Errorf("failed to send post request to %s: %v", r.Config.ApiBaseUrl.Groq, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved invalid status code from server: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (r *Runner) Validate() error {
	if r.ApiKey == "" {
		return ErrEmptyApiKey
	}
	return nil
}
