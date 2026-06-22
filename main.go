package main

import (
	"fmt"
	"github-backup-tool/internal/config"
	"github-backup-tool/internal/git"
	"github-backup-tool/internal/github"
	"github-backup-tool/internal/terminal"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	if cfg.GitHubToken == "" {
		fmt.Println("GITHUB_TOKEN is required. Check your '.env' file")
		return
	}
	fmt.Println("GitHub Token defined successfully!")

	terminal := terminal.NewTerminal(os.Stdin, os.Stdout)

	run(&cfg)

	m, err := terminal.ReadLine("\nMove backup to a snapshot [Y/n]? ")
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(strings.ToLower(m), "n") {
		fmt.Println("Operation canceled.")
		return
	}

	moveBackupToSnapshot(
		"./snapshots/" + time.Now().Format("2006-01-02_15-04-05"),
		"./backup",
	)
}

func run(cfg *config.Config) {
	user, err := github.GetUsername(cfg.GitHubToken)
	if err != nil {
		log.Fatal(err)
	}

	repos, err := github.GetUserRepositories(cfg.GitHubToken)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	for _, repo := range repos {
		if err := cloneRepo(cfg, user, repo); err == nil {
			count++
		}
	}

	fmt.Printf("\nBackup of %v repositories done successfully\n", count)
}

func cloneRepo(
	cfg *config.Config, user, repo string,
) error {
	url := fmt.Sprintf(
		"https://%v:%v@github.com/%v",
		user, cfg.GitHubToken, repo,
	)

	fmt.Println("\nURL to clone:", url)

	repoName := strings.Split(repo, "/")[1]

	clonePath := fmt.Sprintf("./backup/%v", repo)
	mirrorClonePath := fmt.Sprintf("%v/%v.git", clonePath, repoName)
	normalClonePath := fmt.Sprintf("%v/%v", clonePath, repoName)

	if err := git.MirrorClone(url, mirrorClonePath); err != nil {
		log.Print(err, " | skipping...")
		return err
	}
	if err := git.Clone(url, normalClonePath); err != nil {
		log.Print(err, " | skipping...")
		return err
	}

	return nil
}

func moveBackupToSnapshot(dstDir, srcDir string) {
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		log.Fatalf("failed to create destination directory: %s", err)
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatalf("failed to read source directory: %s", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		oldPath := filepath.Join(srcDir, entry.Name())
		newPath := filepath.Join(dstDir, entry.Name())

		// Move the folder
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Fatalf(
				"failed to move %s to %s: %s", 
				entry.Name(), newPath, err,
			)
		}
		fmt.Printf("Moved folder: %s -> %s\n", oldPath, newPath)
	}
}
