package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"errors"
)

func fetchEnv(key string) (string, error) {
	// Open the .env file manually
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error opening .env file:", err)
	}
	defer file.Close()

	// Read file line by line
	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines or comments
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Split key=value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envMap[key] = value
		}
	}

	// Check if key exists
	if token, exists := envMap[key]; exists {
		return token, nil // Return token and nil
	}
	return "", errors.New("key not found") // Return empty string and error if not found
}
