package analysis

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

type result struct {
	name    string
	size_mb float64
}

func Run(path string) {
	results := []result{
		{name: "wasd", size_mb: 25},
		{name: "wasd", size_mb: 25},
	}

	printResult(results)
}

func printResult(results []result) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "Name"},
			{Align: simpletable.AlignLeft, Text: "Size"},
		},
	}

	for _, result := range results {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: result.name},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%.2f MB", result.size_mb)},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
