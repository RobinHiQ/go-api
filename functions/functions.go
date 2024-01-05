package functions

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
)

type Prompt struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func GenerateJobDescription(jobTitle string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := resty.New()

	prompt := Prompt{
		Prompt:    "Skriv en beskrivning på max 535 tecken för jobbet på svenska: " + jobTitle,
		MaxTokens: 200,
	}

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": prompt.Prompt}},
			"max_tokens": prompt.MaxTokens,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error while decoding JSON response:", err)
		return "", err
	}

	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content, nil
}

func GetAllJobs() (string, error) {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	rdb := redis.NewClient(opt)

	var ctx = context.Background()

	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return "", err
	}

	values, err := rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return "", err
	}

	jobs := make(map[string]string)
	for i, key := range keys {
		jobs[key] = values[i].(string)
	}

	jobsJson, err := json.Marshal(jobs)
	if err != nil {
		return "", err
	}

	return string(jobsJson), nil
}
