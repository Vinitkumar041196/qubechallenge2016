package parser

import (
	"encoding/csv"
	"os"
)

func ParseCSVFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return csv.NewReader(file).ReadAll()
}
