package tlsspecloader

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

type CSV struct {
	columns [][]string
	rows    [][]string
}

func ReadCSV(filename string) (*CSV, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse the CSV file
	r := csv.NewReader(bytes.NewReader(fileBytes))
	var rows [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, record)
	}

	// Extract each column as a separate slice
	columns := make([][]string, len(rows[0]))
	for _, row := range rows {
		for i, cell := range row {
			columns[i] = append(columns[i], cell)
		}
	}

	return &CSV{
		columns: columns,
		rows:    rows,
	}, nil
}

func (c *CSV) Column(i int) []string {
	return c.columns[i]
}

func (c *CSV) Row(i int) []string {
	return c.rows[i]
}
