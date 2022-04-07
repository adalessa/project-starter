package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	projectType := fzf([]string{"PHP", "Golang", "Laravel"})
	switch projectType {
	case "PHP":
		createPHP()
	case "Golang":
		createGolang()
	case "Laravel":
		createLaravel()
	default:
		return
	}
}

func checkName(languageFolder string) (string, error) {
	var name string
	fmt.Print("Project Name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return "", fmt.Errorf("Error scanning project name %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting user home dir %w", err)
	}
	projectPath := home + languageFolder + "/" + name
	if _, err := os.Stat(projectPath); err == nil {
		return "", errors.New("Project already exists")
	}
	return projectPath, nil
}
