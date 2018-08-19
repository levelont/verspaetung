package utilities

import (
	"os"
	"log"
	"encoding/csv"
)

func LoadCSVData(pathToCSVFile string) ([][]string) {
	file, err := os.Open(pathToCSVFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return removeHeaderFromRecords(records)
}

func removeHeaderFromRecords(records [][]string) [][]string {
	return records[1:]
}
