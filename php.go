package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const PHP_FOLDER = "/code/php"

func createPHP() {

	projectPath, err := checkName(PHP_FOLDER)
	if err != nil {
		log.Println("Error checking name for create php", err)
		return
	}

	err = os.Mkdir(projectPath, 0755)
	if err != nil {
		log.Println(err)
		return
	}

	os.Chdir(projectPath)
	err = exec.Command("git", "init").Run()
	if err != nil {
		log.Println("Error inizializing git", err)
		return
	}

	composerName := strings.ToLower(os.Getenv("USER") + "/" + filepath.Base(projectPath))
	err = exec.Command("composer", "init", "-n", "--name="+composerName, "-asrc").Run()
	if err != nil {
		log.Println("Error inizializing composer", err)
		return
	}

	fmt.Println(filepath.Base(projectPath))
}
