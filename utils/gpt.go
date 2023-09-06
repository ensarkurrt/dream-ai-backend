package utils

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func GenerateExplanation(prompt string) (string, error) {
	openaiToken := os.Getenv("OPEN_AI_TOKEN")

	client := openai.NewClient(openaiToken)
	ctx := context.Background()

	defaultPrompt := "Sen bir rüya tabir uzmanısın. Sana verilen rüyayı bir uzman gibi yorumla. Rüya:"

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 500,
		Prompt:    defaultPrompt + prompt + "\n\nYorum:",
	}

	resp, err := client.CreateCompletion(ctx, req)

	if err != nil {
		fmt.Printf("Generate Explanation Completion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Text, nil

	/*resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: defaultPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil*/
}

func GenerateTitle(prompt string) (string, error) {

	openaiToken := os.Getenv("OPEN_AI_TOKEN")

	client := openai.NewClient(openaiToken)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 50,
		Prompt:    "Lütfen aşağıdaki metni okuyunuz ve metnin başlığını yazınız.\n\nMetin:\n" + prompt + "\n\nBaşlık:",
	}

	resp, err := client.CreateCompletion(ctx, req)

	if err != nil {
		fmt.Printf("Generate Title Completion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Text, nil
}
