package utils

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func GenerateExplanation(prompt string) (*string, error) {
	openaiToken := os.Getenv("OPEN_AI_TOKEN")

	client := openai.NewClient(openaiToken)
	ctx := context.Background()

	defaultPrompt := "I want you to act as a dream interpreter. \\n  I will give you descriptions of my dreams, and you will provide interpretations based on the symbols and themes present in the dream. \\n  Do not provide personal opinions or assumptions about the dreamer. \\n  Do not summarize the dream. Directly start with interpreting the dream.\\n  Provide only factual interpretations based on the information given. \\n  Write your answer as you are speaking to another person.\\n  Use a friendly tone.: "

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 500,
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
		fmt.Printf("Generate Explanation Completion error: %v\n", err)
		return nil, err
	}

	return &resp.Choices[0].Message.Content, nil
}

func GenerateTitle(prompt string) (*string, error) {

	openaiToken := os.Getenv("OPEN_AI_TOKEN")

	client := openai.NewClient(openaiToken)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 50,
		Prompt:    "Lütfen metin için başlık yazınız.\n\nMetin:\n" + prompt + "\n\nBaşlık:",
	}

	resp, err := client.CreateCompletion(ctx, req)

	if err != nil {
		fmt.Printf("Generate Title Completion error: %v\n", err)
		return nil, err
	}

	return &resp.Choices[0].Text, nil
}

/*
func GenerateImage(prompt string) (*string, error) {
	openaiToken := os.Getenv("OPEN_AI_TOKEN")

	client := openai.NewClient(openaiToken)
	ctx := context.Background()

	imagePrompt := "mdjrny-v4 style there were two loaves of bread. One of them was moldy., digital painting, concept art, smooth, sharp focus, illustration, 8k"

	fmt.Println("Imaginate prompt: " + imagePrompt)

	reqUrl := openai.ImageRequest{
		Prompt:         imagePrompt,
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := client.CreateImage(ctx, reqUrl)

	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return nil, err
	}

	return &respUrl.Data[0].URL, nil
}*/
