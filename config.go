package groq

type Config struct {
	ApiBaseUrl BaseUrlConfig `json:"api_base_url"`
}

type BaseUrlConfig struct {
	Groq string `json:"groq"`
}

func DefaultConfig() Config {
	return Config{
		ApiBaseUrl: BaseUrlConfig{
			Groq: "https://api.groq.com/openai/v1/chat/completions",
		},
	}
}
