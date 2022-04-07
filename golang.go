package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// TODO set this as parameter
const GO_FOLDER = "/code/go/src/github.com/adalessa"

func createGolang() {
	projectPath, err := checkName(GO_FOLDER)
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

	err = exec.Command("go", "mod", "init").Run()
	if err != nil {
		log.Println("Error inizializing go mod", err)
		return
	}

	err = exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		log.Println("Error setting go mod tidy", err)
		return
	}

	fmt.Println(filepath.Base(projectPath))
}
