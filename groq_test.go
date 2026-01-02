package groq

import (
	"os"
	"testing"

	client "github.com/ayuxsec/go-http-client"
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
	clientCfg := client.DefaultClientConfig()
	httpClient, _ := clientCfg.CreateNewClient()
	r := Runner{
		ApiKey:        os.Getenv("GROQ_API_KEY"),
		Client:        httpClient,
		LLMParameters: params,
		Config:        DefaultConfig(),
	}
	resp, err := r.RunGrok()
	if err != nil {
		t.Log(err)
	}
	t.Log(string(resp))
}
