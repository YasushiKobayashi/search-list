package csv_repository

import (
	"encoding/csv"
	"strings"
	"testing"
)

func NewCsvRepository(path string) *CsvRepository {
	return &CsvRepository{
		Path: path,
	}
}

func Test_ReadCsv(t *testing.T) {
	r := NewCsvRepository("sample.csv")
	_, err := r.ReadCsv()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func Test_setCsvInfo(t *testing.T) {
	in := `"LISTING_1","keywords"
"","macbook pro"
"","adwords"`
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvWriter, err := setCsvInfo(records)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if csvWriter.Keywords[0] != "macbook" {
		t.Fatalf("invalid values %#v", csvWriter)
	}
	if csvWriter.Keywords[1] != "adwords" {
		t.Fatalf("invalid values %#v", csvWriter)
	}
}
