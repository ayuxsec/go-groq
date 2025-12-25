package groq

import (
	"net/http"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	client := DefaultClientConfig().CreateNewClient()
	r := Request{
		"https://example.com",
		http.Header{
			"Accept": {"application/json", "text/plain"},
		},
		client,
	}
	resp, err := r.SendPost(strings.NewReader(`{data: "hello server"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	t.Log(resp)
}
