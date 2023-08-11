package actions

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func NavigateDir(current_dir string) (string, error) {
	directories := []string{
		filepath.Dir(current_dir),
	}

	entries, err := os.ReadDir(current_dir)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			directories = append(directories, path.Join(current_dir, entry.Name()))
		}
	}

	for key, dir := range directories {
		fmt.Printf("Key: %d | Directory: %s\n", key, dir)
	}

	fmt.Println()
	fmt.Println("Insert key of directory you want to navigate into")

	var key int
	if _, err := fmt.Scan(&key); err != nil {
		return "", err
	}

	if key < 0 || key > len(directories) {
		return "", errors.New("incorrect directory key was selected")
	}

	return directories[key], nil
}
