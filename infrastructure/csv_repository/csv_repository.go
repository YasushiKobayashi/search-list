package csv_repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

type (
	CsvRepository struct {
		Path string
	}
)

func read(path string) (reader *csv.Reader, file *os.File, err error) {
	file, err = os.Open(path)
	if err != nil {
		return reader, file, errors.Wrap(err, "os.Open error")
	}

	reader = csv.NewReader(file)
	return reader, file, nil
}

// setCsvInfo
// set csv header info models.CsvWriter.Header
// set csv info models.CsvWriter.Keywords
func setCsvInfo(records [][]string) (*model.CsvWriter, error) {
	var csvWriter model.CsvWriter
	var err error
	var keywordsColumnNumber int
	for i, v := range records {
		if i == 0 {
			csvWriter.Header = v
			keywordsColumnNumber, err = getKeywordsColumn(v)
			if err != nil {
				return &csvWriter, errors.Wrap(err, "GetKeywordsColumn error")
			}
		} else {
			csvWriter.Keywords = append(csvWriter.Keywords, model.Keyword(v[keywordsColumnNumber]))
			csvWriter.Rows = append(csvWriter.Rows, v)
		}
	}

	return &csvWriter, nil
}

func getKeywordsColumn(strs []string) (int, error) {
	var errVal int = -1
	var res int = errVal
	const keywords = "keywords"

	for k, v := range strs {
		if v == keywords {
			res = k
			break
		}
	}

	if res == errVal {
		return res, fmt.Errorf("%s not found", keywords)
	}
	return res, nil
}

func (r *CsvRepository) ReadCsv() (csvWriter *model.CsvWriter, err error) {
	reader, file, err := read(r.Path)
	defer file.Close()
	if err != nil {
		return csvWriter, errors.Wrap(err, "read error")
	}

	records, err := reader.ReadAll()
	if err != nil {
		return csvWriter, errors.Wrap(err, "reader.ReadAll error")
	}

	csvWriter, err = setCsvInfo(records)
	if err != nil {
		return csvWriter, errors.Wrap(err, "reader.ReadAll error")
	}
	return csvWriter, nil
}
