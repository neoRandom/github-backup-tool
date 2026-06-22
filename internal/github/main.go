package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetUsername(token string) (string, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		return "", err
	}

	//
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %v", token),
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//
	r, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	//
	var content struct {
		Login string `json:"login"`
	}
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		return "", err
	}

	return content.Login, nil
}

func GetUserRepositories(token string) ([]string, error) {
	repoAmount, err := GetUserRepositoryAmount(token)
	if err != nil {
		return nil, err
	}
	fmt.Println("Amount of repositories: ", repoAmount)

	//
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://api.github.com/user/repos?per_page=%v",
			repoAmount,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	//
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %v", token),
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	//
	var content []struct {
		FullName string `json:"full_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		return nil, err
	}

	//
	repos := make([]string, len(content))
	for i, c := range content {
		repos[i] = c.FullName
	}

	return repos, nil
}

func GetUserRepositoryAmount(token string) (int, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		return 0, err
	}

	//
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %v", token),
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//
	r, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	//
	var content struct {
		PublicRepos       int `json:"public_repos"`
		TotalPrivateRepos int `json:"total_private_repos"`
	}
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		return 0, err
	}

	return content.TotalPrivateRepos + content.PublicRepos, nil
}
