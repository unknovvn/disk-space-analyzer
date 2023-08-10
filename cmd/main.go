package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/unknovvn/disk-space-analyzer/internal/actions"
)

type Action int64

const (
	Retry           Action = 0
	DisplayResults  Action = 1
	NavigateDir     Action = 2
	ExitApplication Action = 9
)

func main() {
	fmt.Println("\nWelcome to Disk Space Analyzer!")

	current_dir := getCurrentDir()
	action := askForAction(current_dir)

	for {
		switch action {
		case Retry:
			action = askForAction(current_dir)
		case DisplayResults:
			actions.DisplayResults(current_dir)
			action = askForAction(current_dir)
		case NavigateDir:
			navigated_dir, err := actions.NavigateDir(current_dir)
			if err == nil {
				current_dir = navigated_dir
			} else {
				log.Printf("Wasn't able to navigate directory. Error: %s", err)
			}

			action = askForAction(current_dir)
		case ExitApplication:
			os.Exit(0)
		}
	}
}

func askForAction(current_dir string) Action {
	fmt.Printf("\nDirectory you're currently in: %s", current_dir)
	fmt.Print("\nChoose your action (type action number):")
	fmt.Println()
	fmt.Printf("%v - Display analysis results \n", DisplayResults)
	fmt.Printf("%v - Navigate directory \n", NavigateDir)
	fmt.Println()
	fmt.Printf("%v - Exit app \n", ExitApplication)
	fmt.Println()

	var action Action
	if _, err := fmt.Scan(&action); err != nil {
		fmt.Printf("Error occured while selecting action: %s", err)

		return Retry
	}

	return action
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Dir(ex)
}
