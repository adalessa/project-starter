package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func createLaravel() {
	projectPath, err := checkName(PHP_FOLDER)
	if err != nil {
		log.Println("Error checking name for create php", err)
		return
	}

	os.Chdir(filepath.Dir(projectPath))
	cmd := fmt.Sprintf("curl -s https://laravel.build/%s | bash", filepath.Base(projectPath))
	err = exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("Error when calling laravel creation", err)
		return
	}
}
