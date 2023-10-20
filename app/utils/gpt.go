package utils

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

type GPT interface {
	GenerateExplanation(prompt string) (string, error)
	GenerateTitle(prompt string) (string, error)
	GenerateImagePrompt(prompt string) (string, error)
}

type GPTImpl struct {
	client *openai.Client
	ctx    context.Context
}

func NewGPT() *GPTImpl {
	openaiToken := os.Getenv("OPEN_AI_TOKEN")

	client := openai.NewClient(openaiToken)
	ctx := context.Background()

	return &GPTImpl{
		client: client,
		ctx:    ctx,
	}
}

func (gpt *GPTImpl) GenerateExplanation(prompt string) (string, error) {
	defaultPrompt := "Rüya yorumcusu olmanı istiyorum. \\n Ben size rüyalarımın tanımlarını vereceğim, siz de rüyada mevcut olan sembol ve temalara göre yorumlar yapacaksınız. \\n Rüyayı gören kişi hakkında kişisel görüş veya varsayımlarda bulunmayın. \\n Rüyayı özetlemeyin. Doğrudan rüyayı yorumlamakla başlayın.\\n Verilen bilgilere dayanarak yalnızca gerçeklere dayanan yorumlar yapın. \\n Cevabınızı başka biriyle konuşurken yazın.\\n Dostça bir ses tonu kullanın.: "

	resp, err := gpt.client.CreateChatCompletion(
		gpt.ctx,
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

	log.Println("Prompt: ", prompt)
	if err != nil {
		log.Println("Generate Explanation Completion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (gpt *GPTImpl) GenerateTitle(prompt string) (string, error) {

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 50,
		Prompt:    "Lütfen metin için başlık yazınız.\n\nMetin:\n" + prompt + "\n\nBaşlık:",
	}

	resp, err := gpt.client.CreateCompletion(gpt.ctx, req)

	if err != nil {
		log.Println("Generate Title Completion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Text, nil
}

func (gpt *GPTImpl) GenerateImagePrompt(prompt string) (string, error) {
	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 50,
		Prompt:    " \" " + prompt + " \" metninin ingilizcesi:",
	}

	resp, err := gpt.client.CreateCompletion(gpt.ctx, req)

	if err != nil {
		log.Println("Generate Image Prompt Completion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Text, nil
}
