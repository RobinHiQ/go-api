package functions

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

type Prompt struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

func GenerateJobDescription(jobTitle string) (string, error) {
	prompt := Prompt{
		Prompt:    "Generate a job description for a " + jobTitle,
		MaxTokens: 200,
	}

	body, err := json.Marshal(prompt)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci-codex/completions", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	description := result["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)

	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	rdb := redis.NewClient(opt)

	var ctx = context.Background()

	err = rdb.Set(ctx, jobTitle, description, 0).Err()
	if err != nil {
		return "", err
	}

	return description, nil
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
