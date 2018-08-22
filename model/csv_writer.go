package model

type (
	CsvWriter struct {
		Header   []string
		Rows     [][]string
		Keywords []Keyword
		SearchLists
	}
)
