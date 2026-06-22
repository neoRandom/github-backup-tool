package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Config holds environment values loaded from a local .env file.
type Config struct {
	GitHubToken string
}

// LoadConfig reads key/value pairs from the given .env file path.
func LoadConfig(path string) (Config, error) {
	values, err := readDotEnv(path)
	if err != nil {
		return Config{}, err
	}

	return Config{
		GitHubToken: values["GITHUB_TOKEN"],
	}, nil
}

func readDotEnv(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	values := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			return nil, fmt.Errorf("invalid .env line: %q", line)
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" {
			return nil, fmt.Errorf("invalid .env line: %q", line)
		}

		values[key] = trimQuotes(value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

func trimQuotes(value string) string {
	if len(value) >= 2 {
		first := value[0]
		last := value[len(value)-1]
		if (first == '"' && last == '"') || (first == '\'' && last == '\'') {
			return value[1 : len(value)-1]
		}
	}

	return value
}
