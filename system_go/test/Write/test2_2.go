package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		{"firstname", "lastname", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Harutugu", "Iwata", "harukasu"},
	}

	file, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(os.Stdout)
	fw := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			panic(err)
		}
		if err := fw.Write(record); err != nil {
			panic(err)
		}
	}

	w.Flush()
	fw.Flush()
}
