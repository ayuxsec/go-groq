package groq

import (
	"os"
	"testing"
)

func TestRunGrok(t *testing.T) {
	params := DefaultLLMParameters()
	params.Model = "openai/gpt-oss-120b"
	params.LLMMessages = []Message{
		{
			Role:    "user",
			Content: "Ping",
		},
	}
	r := Runner{
		ApiKey:        os.Getenv("GROQ_API_KEY"),
		Client:        DefaultClientConfig().CreateNewClient(),
		LLMParameters: params,
		Config:        DefaultConfig(),
	}
	resp, err := r.RunGrok()
	if err != nil {
		t.Log(err)
	}
	t.Log(string(resp))
}
