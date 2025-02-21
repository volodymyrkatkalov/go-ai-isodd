package isodd

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	client *openai.Client
}

func NewChatGPTClient(apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		client: openai.NewClient(apiKey),
	}
}

func (c *ChatGPTClient) IsOdd(number int) (bool, error) {
	prompt := fmt.Sprintf("Is %d an odd number? Answer with 'yes' or 'no' only.", number)
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 10, // Limit response length
		},
	)
	if err != nil {
		return false, fmt.Errorf("failed to get response from ChatGPT: %v", err)
	}

	if len(resp.Choices) == 0 {
		return false, fmt.Errorf("no response from ChatGPT")
	}

	answer := strings.ToLower(strings.TrimSpace(resp.Choices[0].Message.Content))
	fmt.Printf("ChatGPT says: %d is %s\n", number, answer)

	switch answer {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	default:
		return false, fmt.Errorf("unexpected response from ChatGPT: %s", answer)
	}
}
