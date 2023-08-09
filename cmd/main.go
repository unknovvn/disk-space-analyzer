package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/unknovvn/disk-space-analyzer/internal/analysis"
)

func main() {
	fmt.Println("Disk Space Analyzer")

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	analysis.Run(exPath)
}
