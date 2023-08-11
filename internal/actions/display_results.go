package actions

import (
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/unknovvn/disk-space-analyzer/internal/analysis"
)

func DisplayResults(current_dir string) {
	entries, err := os.ReadDir(current_dir)
	if err != nil {
		displayError(err)
		return
	}

	results, err := analysis.Run(current_dir, entries)
	if err != nil {
		displayError(err)
		return
	}

	printResult(results)
}

func displayError(err error) {
	fmt.Printf("Error occured while trying to display results: %s", err)
}

func printResult(results []analysis.Record) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "Name"},
			{Align: simpletable.AlignLeft, Text: "Size"},
		},
	}

	for _, result := range results {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: result.Name},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d bytes", result.Size_bytes)},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleCompactLite)

	fmt.Println("Analysis results:")
	fmt.Println()
	fmt.Println(table.String())
}
