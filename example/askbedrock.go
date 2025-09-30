package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"io"
	"net/http"
	"os"
)

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/aws/aws-sdk-go-v2/aws"
//	"github.com/aws/aws-sdk-go-v2/config"
//	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
//"context"
//"github.com/aws/aws-sdk-go-v2/config"
//"github.com/aws/aws-sdk-go-v2/service/polly"
//"github.com/aws/aws-sdk-go-v2/service/polly/types"
//"io/ioutil"
//)
//
//func AskBedrock(question string) (string, error) {
//	cfg, err := config.LoadDefaultConfig(context.TODO())
//	if err != nil {
//		return "", err
//	}
//	client := bedrockruntime.NewFromConfig(cfg)
//
//	payload := map[string]interface{}{
//		"prompt":               question,
//		"max_tokens_to_sample": 256,
//	}
//	body, _ := json.Marshal(payload)
//
//	input := &bedrockruntime.InvokeModelInput{
//		ModelId:     aws.String("amazon.titan-text-express-v1"),
//		ContentType: aws.String("application/json"),
//		Body:        body,
//	}
//	resp, err := client.InvokeModel(context.TODO(), input)
//	if err != nil {
//		return "", err
//	}
//
//	var result struct {
//		Completion string `json:"completion"`
//	}
//	if err := json.Unmarshal(resp.Body, &result); err != nil {
//		return "", fmt.Errorf("解析Bedrock返回内容失败: %v", err)
//	}
//	return result.Completion, nil
//}

func AskDeepSeek(messages []map[string]string) (string, error) {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("请先设置 DEEPSEEK_API_KEY 环境变量")
	}
	url := "https://api.deepseek.com/v1/chat/completions"
	payload := map[string]interface{}{
		"model":      "deepseek-chat",
		"messages":   messages,
		"max_tokens": 256,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("解析DeepSeek返回内容失败: %v", err)
	}
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("DeepSeek无返回内容")
	}
	return result.Choices[0].Message.Content, nil
}

func SynthesizeSpeech(text, filename string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := polly.NewFromConfig(cfg)
	input := &polly.SynthesizeSpeechInput{
		OutputFormat: types.OutputFormatMp3,
		Text:         &text,
		VoiceId:      types.VoiceIdZhiyu,
		LanguageCode: "cmn-CN",
	}
	resp, err := client.SynthesizeSpeech(context.TODO(), input)
	if err != nil {
		return err
	}
	defer resp.AudioStream.Close()
	audio, err := io.ReadAll(resp.AudioStream)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, audio, 0644)
}
