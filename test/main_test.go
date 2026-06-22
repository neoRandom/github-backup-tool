package test

import (
	"fmt"
	"github-backup-tool/internal/config"
	"github-backup-tool/internal/github"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
	}
	if cfg.GitHubToken == "" {
		log.Fatal("No GITHUB_TOKEN defined")
	}

	repos, err := github.GetUserRepositories(cfg.GitHubToken)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range repos {
		fmt.Println(r)
	}

	fmt.Println("\nNumber of repositories: ", len(repos))
}
