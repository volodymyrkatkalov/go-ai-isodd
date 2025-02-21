package isodd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sashabaranov/go-openai"
)

// Mock server to simulate OpenAI API
func setupMockServer(t *testing.T) (*httptest.Server, *ChatGPTClient) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Simple mock: assume any number ending in 0,2,4,6,8 is even
		body := struct {
			Messages []struct {
				Content string `json:"content"`
			} `json:"messages"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if len(body.Messages) == 0 {
			http.Error(w, "no messages", http.StatusBadRequest)
			return
		}
		var num int
		fmt.Sscanf(body.Messages[0].Content, "Is %d an odd number? Answer with 'yes' or 'no' only.", &num)
		isOdd := num%2 != 0
		response := `{"choices":[{"message":{"content":"` + map[bool]string{true: "yes", false: "no"}[isOdd] + `"}}]}`
		w.Write([]byte(response))
	}))

	config := openai.DefaultConfig("test-key")
	config.BaseURL = server.URL
	client := openai.NewClientWithConfig(config)
	return server, &ChatGPTClient{client: client}
}

func TestIsOdd(t *testing.T) {
	server, client := setupMockServer(t)
	defer server.Close()

	tests := []struct {
		input    int
		expected bool
	}{
		{2, false},
		{3, true},
		{0, false},
		{-1, true},
		{-2, false},
	}

	for _, test := range tests {
		result, err := client.IsOdd(test.input)
		if err != nil {
			t.Errorf("IsOdd(%d) returned error: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("IsOdd(%d) = %v, want %v", test.input, result, test.expected)
		}
	}
}
